// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomFieldOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetCustomFieldOptionResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetCustomFieldOptionResponseBody
	GetErrorMsg() *string
	SetFileds(v []*GetCustomFieldOptionResponseBodyFileds) *GetCustomFieldOptionResponseBody
	GetFileds() []*GetCustomFieldOptionResponseBodyFileds
	SetRequestId(v string) *GetCustomFieldOptionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCustomFieldOptionResponseBody
	GetSuccess() *bool
}

type GetCustomFieldOptionResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string                                   `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Fileds   []*GetCustomFieldOptionResponseBodyFileds `json:"fileds,omitempty" xml:"fileds,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetCustomFieldOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldOptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetCustomFieldOptionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetCustomFieldOptionResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetCustomFieldOptionResponseBody) GetFileds() []*GetCustomFieldOptionResponseBodyFileds {
	return s.Fileds
}

func (s *GetCustomFieldOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCustomFieldOptionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCustomFieldOptionResponseBody) SetErrorCode(v string) *GetCustomFieldOptionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetCustomFieldOptionResponseBody) SetErrorMsg(v string) *GetCustomFieldOptionResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetCustomFieldOptionResponseBody) SetFileds(v []*GetCustomFieldOptionResponseBodyFileds) *GetCustomFieldOptionResponseBody {
	s.Fileds = v
	return s
}

func (s *GetCustomFieldOptionResponseBody) SetRequestId(v string) *GetCustomFieldOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCustomFieldOptionResponseBody) SetSuccess(v bool) *GetCustomFieldOptionResponseBody {
	s.Success = &v
	return s
}

func (s *GetCustomFieldOptionResponseBody) Validate() error {
	if s.Fileds != nil {
		for _, item := range s.Fileds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCustomFieldOptionResponseBodyFileds struct {
	// example:
	//
	// 223
	DisplayValue *string `json:"displayValue,omitempty" xml:"displayValue,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	FieldIdentifier *string `json:"fieldIdentifier,omitempty" xml:"fieldIdentifier,omitempty"`
	// example:
	//
	// 3345
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 1
	Level *int64 `json:"level,omitempty" xml:"level,omitempty"`
	// example:
	//
	// 1
	Position *int64 `json:"position,omitempty" xml:"position,omitempty"`
	// example:
	//
	// 223
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// 223
	ValueEn *string `json:"valueEn,omitempty" xml:"valueEn,omitempty"`
}

func (s GetCustomFieldOptionResponseBodyFileds) String() string {
	return dara.Prettify(s)
}

func (s GetCustomFieldOptionResponseBodyFileds) GoString() string {
	return s.String()
}

func (s *GetCustomFieldOptionResponseBodyFileds) GetDisplayValue() *string {
	return s.DisplayValue
}

func (s *GetCustomFieldOptionResponseBodyFileds) GetFieldIdentifier() *string {
	return s.FieldIdentifier
}

func (s *GetCustomFieldOptionResponseBodyFileds) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetCustomFieldOptionResponseBodyFileds) GetLevel() *int64 {
	return s.Level
}

func (s *GetCustomFieldOptionResponseBodyFileds) GetPosition() *int64 {
	return s.Position
}

func (s *GetCustomFieldOptionResponseBodyFileds) GetValue() *string {
	return s.Value
}

func (s *GetCustomFieldOptionResponseBodyFileds) GetValueEn() *string {
	return s.ValueEn
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetDisplayValue(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.DisplayValue = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetFieldIdentifier(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.FieldIdentifier = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetIdentifier(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.Identifier = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetLevel(v int64) *GetCustomFieldOptionResponseBodyFileds {
	s.Level = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetPosition(v int64) *GetCustomFieldOptionResponseBodyFileds {
	s.Position = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetValue(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.Value = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) SetValueEn(v string) *GetCustomFieldOptionResponseBodyFileds {
	s.ValueEn = &v
	return s
}

func (s *GetCustomFieldOptionResponseBodyFileds) Validate() error {
	return dara.Validate(s)
}
