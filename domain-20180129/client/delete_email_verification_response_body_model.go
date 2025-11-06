// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmailVerificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailList(v []*DeleteEmailVerificationResponseBodyFailList) *DeleteEmailVerificationResponseBody
	GetFailList() []*DeleteEmailVerificationResponseBodyFailList
	SetRequestId(v string) *DeleteEmailVerificationResponseBody
	GetRequestId() *string
	SetSuccessList(v []*DeleteEmailVerificationResponseBodySuccessList) *DeleteEmailVerificationResponseBody
	GetSuccessList() []*DeleteEmailVerificationResponseBodySuccessList
}

type DeleteEmailVerificationResponseBody struct {
	FailList []*DeleteEmailVerificationResponseBodyFailList `json:"FailList,omitempty" xml:"FailList,omitempty" type:"Repeated"`
	// example:
	//
	// 7A3D0E4A-0D4B-4BD0-90D7-A61DF8DD26AE
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessList []*DeleteEmailVerificationResponseBodySuccessList `json:"SuccessList,omitempty" xml:"SuccessList,omitempty" type:"Repeated"`
}

func (s DeleteEmailVerificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmailVerificationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationResponseBody) GetFailList() []*DeleteEmailVerificationResponseBodyFailList {
	return s.FailList
}

func (s *DeleteEmailVerificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEmailVerificationResponseBody) GetSuccessList() []*DeleteEmailVerificationResponseBodySuccessList {
	return s.SuccessList
}

func (s *DeleteEmailVerificationResponseBody) SetFailList(v []*DeleteEmailVerificationResponseBodyFailList) *DeleteEmailVerificationResponseBody {
	s.FailList = v
	return s
}

func (s *DeleteEmailVerificationResponseBody) SetRequestId(v string) *DeleteEmailVerificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEmailVerificationResponseBody) SetSuccessList(v []*DeleteEmailVerificationResponseBodySuccessList) *DeleteEmailVerificationResponseBody {
	s.SuccessList = v
	return s
}

func (s *DeleteEmailVerificationResponseBody) Validate() error {
	if s.FailList != nil {
		for _, item := range s.FailList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessList != nil {
		for _, item := range s.SuccessList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteEmailVerificationResponseBodyFailList struct {
	// example:
	//
	// ParameterIllegall
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test1@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// Parameter error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteEmailVerificationResponseBodyFailList) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmailVerificationResponseBodyFailList) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationResponseBodyFailList) GetCode() *string {
	return s.Code
}

func (s *DeleteEmailVerificationResponseBodyFailList) GetEmail() *string {
	return s.Email
}

func (s *DeleteEmailVerificationResponseBodyFailList) GetMessage() *string {
	return s.Message
}

func (s *DeleteEmailVerificationResponseBodyFailList) SetCode(v string) *DeleteEmailVerificationResponseBodyFailList {
	s.Code = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodyFailList) SetEmail(v string) *DeleteEmailVerificationResponseBodyFailList {
	s.Email = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodyFailList) SetMessage(v string) *DeleteEmailVerificationResponseBodyFailList {
	s.Message = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodyFailList) Validate() error {
	return dara.Validate(s)
}

type DeleteEmailVerificationResponseBodySuccessList struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// test2@aliyun.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteEmailVerificationResponseBodySuccessList) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmailVerificationResponseBodySuccessList) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationResponseBodySuccessList) GetCode() *string {
	return s.Code
}

func (s *DeleteEmailVerificationResponseBodySuccessList) GetEmail() *string {
	return s.Email
}

func (s *DeleteEmailVerificationResponseBodySuccessList) GetMessage() *string {
	return s.Message
}

func (s *DeleteEmailVerificationResponseBodySuccessList) SetCode(v string) *DeleteEmailVerificationResponseBodySuccessList {
	s.Code = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodySuccessList) SetEmail(v string) *DeleteEmailVerificationResponseBodySuccessList {
	s.Email = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodySuccessList) SetMessage(v string) *DeleteEmailVerificationResponseBodySuccessList {
	s.Message = &v
	return s
}

func (s *DeleteEmailVerificationResponseBodySuccessList) Validate() error {
	return dara.Validate(s)
}
