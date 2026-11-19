// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePersonalICodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdatePersonalICodeResponseBody
	GetRequestId() *string
	SetResult(v *UpdatePersonalICodeResponseBodyResult) *UpdatePersonalICodeResponseBody
	GetResult() *UpdatePersonalICodeResponseBodyResult
	SetSuccess(v bool) *UpdatePersonalICodeResponseBody
	GetSuccess() *bool
}

type UpdatePersonalICodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 685072a0-1fd5-40ef-ae6b-cf94e79e718f
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The AK model.
	//
	// example:
	//
	// {
	//
	//         "icode": "V2Yc/dMFWBwTKLNCNxgnN****z***5ARu",
	//
	//         "icodeKey": "b6f3***32a62936",
	//
	//         "type": 1
	//
	//     }
	Result *UpdatePersonalICodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePersonalICodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePersonalICodeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePersonalICodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePersonalICodeResponseBody) GetResult() *UpdatePersonalICodeResponseBodyResult {
	return s.Result
}

func (s *UpdatePersonalICodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdatePersonalICodeResponseBody) SetRequestId(v string) *UpdatePersonalICodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePersonalICodeResponseBody) SetResult(v *UpdatePersonalICodeResponseBodyResult) *UpdatePersonalICodeResponseBody {
	s.Result = v
	return s
}

func (s *UpdatePersonalICodeResponseBody) SetSuccess(v bool) *UpdatePersonalICodeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdatePersonalICodeResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdatePersonalICodeResponseBodyResult struct {
	// SK
	//
	// example:
	//
	// V2Yc/dMFWBwTKLNCNxgnN****z***5ARu
	Icode *string `json:"Icode,omitempty" xml:"Icode,omitempty"`
	// AK
	//
	// example:
	//
	// b6f3***32a62936
	IcodeKey *string `json:"IcodeKey,omitempty" xml:"IcodeKey,omitempty"`
	// The AK type. Currently, only the value 1 is supported, which indicates a personal AK.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePersonalICodeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpdatePersonalICodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdatePersonalICodeResponseBodyResult) GetIcode() *string {
	return s.Icode
}

func (s *UpdatePersonalICodeResponseBodyResult) GetIcodeKey() *string {
	return s.IcodeKey
}

func (s *UpdatePersonalICodeResponseBodyResult) GetType() *int32 {
	return s.Type
}

func (s *UpdatePersonalICodeResponseBodyResult) SetIcode(v string) *UpdatePersonalICodeResponseBodyResult {
	s.Icode = &v
	return s
}

func (s *UpdatePersonalICodeResponseBodyResult) SetIcodeKey(v string) *UpdatePersonalICodeResponseBodyResult {
	s.IcodeKey = &v
	return s
}

func (s *UpdatePersonalICodeResponseBodyResult) SetType(v int32) *UpdatePersonalICodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *UpdatePersonalICodeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
