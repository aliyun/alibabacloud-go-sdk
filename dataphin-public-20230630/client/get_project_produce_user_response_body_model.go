// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectProduceUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetProjectProduceUserResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetProjectProduceUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetProjectProduceUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetProjectProduceUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProjectProduceUserResponseBody
	GetSuccess() *bool
	SetUser(v *GetProjectProduceUserResponseBodyUser) *GetProjectProduceUserResponseBody
	GetUser() *GetProjectProduceUserResponseBodyUser
}

type GetProjectProduceUserResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	User    *GetProjectProduceUserResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetProjectProduceUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectProduceUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectProduceUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetProjectProduceUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetProjectProduceUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetProjectProduceUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectProduceUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectProduceUserResponseBody) GetUser() *GetProjectProduceUserResponseBodyUser {
	return s.User
}

func (s *GetProjectProduceUserResponseBody) SetCode(v string) *GetProjectProduceUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetHttpStatusCode(v int32) *GetProjectProduceUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetMessage(v string) *GetProjectProduceUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetRequestId(v string) *GetProjectProduceUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetSuccess(v bool) *GetProjectProduceUserResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectProduceUserResponseBody) SetUser(v *GetProjectProduceUserResponseBodyUser) *GetProjectProduceUserResponseBody {
	s.User = v
	return s
}

func (s *GetProjectProduceUserResponseBody) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectProduceUserResponseBodyUser struct {
	// example:
	//
	// 123131
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProjectProduceUserResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s GetProjectProduceUserResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetProjectProduceUserResponseBodyUser) GetId() *string {
	return s.Id
}

func (s *GetProjectProduceUserResponseBodyUser) SetId(v string) *GetProjectProduceUserResponseBodyUser {
	s.Id = &v
	return s
}

func (s *GetProjectProduceUserResponseBodyUser) Validate() error {
	return dara.Validate(s)
}
