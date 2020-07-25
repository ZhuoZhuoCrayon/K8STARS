// Package TestApp comment
// This file was generated by tars2go 1.1.4
// Generated from SayHello.tars
package TestApp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/basef"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/protocol/tup"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
	"unsafe"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = codec.FromInt8
var _ = unsafe.Pointer(nil)

//SayHello struct
type SayHello struct {
	s m.Servant
}

//Add is the proxy function for the method defined in the tars file, with the context
func (_obj *SayHello) Add(a int32, b int32, c *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(a, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(b, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32((*c), 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "Add", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*c), 3, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//AddWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *SayHello) AddWithContext(tarsCtx context.Context, a int32, b int32, c *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(a, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(b, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32((*c), 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "Add", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*c), 3, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//AddOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *SayHello) AddOneWayWithContext(tarsCtx context.Context, a int32, b int32, c *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(a, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(b, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32((*c), 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "Add", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Sub is the proxy function for the method defined in the tars file, with the context
func (_obj *SayHello) Sub(a int32, b int32, c *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(a, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(b, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32((*c), 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	tarsCtx := context.Background()

	err = _obj.s.Tars_invoke(tarsCtx, 0, "Sub", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*c), 3, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SubWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *SayHello) SubWithContext(tarsCtx context.Context, a int32, b int32, c *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(a, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(b, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32((*c), 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 0, "Sub", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*c), 3, true)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SubOneWayWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *SayHello) SubOneWayWithContext(tarsCtx context.Context, a int32, b int32, c *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_int32(a, 1)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32(b, 2)
	if err != nil {
		return ret, err
	}

	err = _os.Write_int32((*c), 3)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)

	err = _obj.s.Tars_invoke(tarsCtx, 1, "Sub", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range _resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range _resp.Status {
			_status[k] = v
		}

	}
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *SayHello) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *SayHello) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol sets the protocol for the servant.
func (_obj *SayHello) TarsSetProtocol(p m.Protocol) {
	_obj.s.TarsSetProtocol(p)
}

//AddServant adds servant  for the service.
func (_obj *SayHello) AddServant(imp _impSayHello, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *SayHello) AddServantWithContext(imp _impSayHelloWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impSayHello interface {
	Add(a int32, b int32, c *int32) (ret int32, err error)
	Sub(a int32, b int32, c *int32) (ret int32, err error)
}
type _impSayHelloWithContext interface {
	Add(tarsCtx context.Context, a int32, b int32, c *int32) (ret int32, err error)
	Sub(tarsCtx context.Context, a int32, b int32, c *int32) (ret int32, err error)
}

// Dispatch is used to call the server side implemnet for the method defined in the tars file. _withContext shows using context or not.
func (_obj *SayHello) Dispatch(tarsCtx context.Context, _val interface{}, tarsReq *requestf.RequestPacket, tarsResp *requestf.ResponsePacket, _withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	_is := codec.NewReader(tools.Int8ToByte(tarsReq.SBuffer))
	_os := codec.NewBuffer()
	switch tarsReq.SFuncName {
	case "Add":
		var a int32
		var b int32
		var c int32

		if tarsReq.IVersion == basef.TARSVERSION {

			err = _is.Read_int32(&a, 1, true)
			if err != nil {
				return err
			}

			err = _is.Read_int32(&b, 2, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("a", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = _is.Read_int32(&a, 0, true)
			if err != nil {
				return err
			}

			_reqTup_.GetBuffer("b", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = _is.Read_int32(&b, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["a"])
				if err = json.Unmarshal([]byte(_jsonStr_), &a); err != nil {
					return err
				}
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["b"])
				if err = json.Unmarshal([]byte(_jsonStr_), &b); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impSayHello)
			_funRet_, err = _imp.Add(a, b, &c)
		} else {
			_imp := _val.(_impSayHelloWithContext)
			_funRet_, err = _imp.Add(tarsCtx, a, b, &c)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = _os.Write_int32(c, 3)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = _os.Write_int32(c, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("c", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_
			_rspJson_["c"] = c

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}
	case "Sub":
		var a int32
		var b int32
		var c int32

		if tarsReq.IVersion == basef.TARSVERSION {

			err = _is.Read_int32(&a, 1, true)
			if err != nil {
				return err
			}

			err = _is.Read_int32(&b, 2, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_reqTup_ := tup.NewUniAttribute()
			_reqTup_.Decode(_is)

			var _tupBuffer_ []byte

			_reqTup_.GetBuffer("a", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = _is.Read_int32(&a, 0, true)
			if err != nil {
				return err
			}

			_reqTup_.GetBuffer("b", &_tupBuffer_)
			_is.Reset(_tupBuffer_)
			err = _is.Read_int32(&b, 0, true)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.JSONVERSION {
			var _jsonDat_ map[string]interface{}
			err = json.Unmarshal(_is.ToBytes(), &_jsonDat_)
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["a"])
				if err = json.Unmarshal([]byte(_jsonStr_), &a); err != nil {
					return err
				}
			}
			{
				_jsonStr_, _ := json.Marshal(_jsonDat_["b"])
				if err = json.Unmarshal([]byte(_jsonStr_), &b); err != nil {
					return err
				}
			}

		} else {
			err = fmt.Errorf("Decode reqpacket fail, error version:", tarsReq.IVersion)
			return err
		}

		var _funRet_ int32
		if _withContext == false {
			_imp := _val.(_impSayHello)
			_funRet_, err = _imp.Sub(a, b, &c)
		} else {
			_imp := _val.(_impSayHelloWithContext)
			_funRet_, err = _imp.Sub(tarsCtx, a, b, &c)
		}

		if err != nil {
			return err
		}

		if tarsReq.IVersion == basef.TARSVERSION {
			_os.Reset()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			err = _os.Write_int32(c, 3)
			if err != nil {
				return err
			}

		} else if tarsReq.IVersion == basef.TUPVERSION {
			_tupRsp_ := tup.NewUniAttribute()

			err = _os.Write_int32(_funRet_, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("", _os.ToBytes())
			_tupRsp_.PutBuffer("tars_ret", _os.ToBytes())

			_os.Reset()
			err = _os.Write_int32(c, 0)
			if err != nil {
				return err
			}

			_tupRsp_.PutBuffer("c", _os.ToBytes())

			_os.Reset()
			err = _tupRsp_.Encode(_os)
			if err != nil {
				return err
			}
		} else if tarsReq.IVersion == basef.JSONVERSION {
			_rspJson_ := map[string]interface{}{}
			_rspJson_["tars_ret"] = _funRet_
			_rspJson_["c"] = c

			var _rspByte_ []byte
			if _rspByte_, err = json.Marshal(_rspJson_); err != nil {
				return err
			}

			_os.Reset()
			err = _os.Write_slice_uint8(_rspByte_)
			if err != nil {
				return err
			}
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(tarsCtx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(tarsCtx)
	if ok && c != nil {
		_context = c
	}
	*tarsResp = requestf.ResponsePacket{
		IVersion:     tarsReq.IVersion,
		CPacketType:  0,
		IRequestId:   tarsReq.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}

	_ = _is
	_ = _os
	_ = length
	_ = have
	_ = ty
	return nil
}
