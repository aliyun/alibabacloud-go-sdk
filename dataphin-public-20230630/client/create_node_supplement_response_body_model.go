// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeSupplementResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateNodeSupplementResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateNodeSupplementResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateNodeSupplementResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateNodeSupplementResponseBody
	GetRequestId() *string
	SetSubmitId(v string) *CreateNodeSupplementResponseBody
	GetSubmitId() *string
	SetSuccess(v bool) *CreateNodeSupplementResponseBody
	GetSuccess() *bool
}

type CreateNodeSupplementResponseBody struct {
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
	// f_2264518792396800000_20210223_2329354897145659392
	SubmitId *string `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
	// example:
	//
	// true/false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNodeSupplementResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeSupplementResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeSupplementResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateNodeSupplementResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateNodeSupplementResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateNodeSupplementResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodeSupplementResponseBody) GetSubmitId() *string {
	return s.SubmitId
}

func (s *CreateNodeSupplementResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNodeSupplementResponseBody) SetCode(v string) *CreateNodeSupplementResponseBody {
	s.Code = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetHttpStatusCode(v int32) *CreateNodeSupplementResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetMessage(v string) *CreateNodeSupplementResponseBody {
	s.Message = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetRequestId(v string) *CreateNodeSupplementResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetSubmitId(v string) *CreateNodeSupplementResponseBody {
	s.SubmitId = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) SetSuccess(v bool) *CreateNodeSupplementResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNodeSupplementResponseBody) Validate() error {
	return dara.Validate(s)
}
