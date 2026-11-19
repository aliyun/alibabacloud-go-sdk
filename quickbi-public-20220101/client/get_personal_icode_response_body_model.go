// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPersonalICodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetPersonalICodeResponseBody
	GetRequestId() *string
	SetResult(v *GetPersonalICodeResponseBodyResult) *GetPersonalICodeResponseBody
	GetResult() *GetPersonalICodeResponseBodyResult
	SetSuccess(v bool) *GetPersonalICodeResponseBody
	GetSuccess() *bool
}

type GetPersonalICodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BCE45E6D-9304-4F94-86BB-5A772B1615FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The AccessKey model.
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
	Result *GetPersonalICodeResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
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

func (s GetPersonalICodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPersonalICodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetPersonalICodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPersonalICodeResponseBody) GetResult() *GetPersonalICodeResponseBodyResult {
	return s.Result
}

func (s *GetPersonalICodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPersonalICodeResponseBody) SetRequestId(v string) *GetPersonalICodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPersonalICodeResponseBody) SetResult(v *GetPersonalICodeResponseBodyResult) *GetPersonalICodeResponseBody {
	s.Result = v
	return s
}

func (s *GetPersonalICodeResponseBody) SetSuccess(v bool) *GetPersonalICodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetPersonalICodeResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPersonalICodeResponseBodyResult struct {
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
	// asdasd
	IcodeKey *string `json:"IcodeKey,omitempty" xml:"IcodeKey,omitempty"`
	// The AccessKey type. Currently, only the value 1 is supported, which indicates a personal AccessKey.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetPersonalICodeResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetPersonalICodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetPersonalICodeResponseBodyResult) GetIcode() *string {
	return s.Icode
}

func (s *GetPersonalICodeResponseBodyResult) GetIcodeKey() *string {
	return s.IcodeKey
}

func (s *GetPersonalICodeResponseBodyResult) GetType() *int32 {
	return s.Type
}

func (s *GetPersonalICodeResponseBodyResult) SetIcode(v string) *GetPersonalICodeResponseBodyResult {
	s.Icode = &v
	return s
}

func (s *GetPersonalICodeResponseBodyResult) SetIcodeKey(v string) *GetPersonalICodeResponseBodyResult {
	s.IcodeKey = &v
	return s
}

func (s *GetPersonalICodeResponseBodyResult) SetType(v int32) *GetPersonalICodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetPersonalICodeResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
