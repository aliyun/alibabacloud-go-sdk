// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGuideSubStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryGuideSubStatusResponseBody
	GetCode() *string
	SetData(v string) *QueryGuideSubStatusResponseBody
	GetData() *string
	SetMessage(v string) *QueryGuideSubStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryGuideSubStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryGuideSubStatusResponseBody
	GetSuccess() *bool
}

type QueryGuideSubStatusResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// SUBSCRIBED
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 3C107939-59BD-5EB9-B250-39559C830A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryGuideSubStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGuideSubStatusResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGuideSubStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryGuideSubStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryGuideSubStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryGuideSubStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGuideSubStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryGuideSubStatusResponseBody) SetCode(v string) *QueryGuideSubStatusResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGuideSubStatusResponseBody) SetData(v string) *QueryGuideSubStatusResponseBody {
	s.Data = &v
	return s
}

func (s *QueryGuideSubStatusResponseBody) SetMessage(v string) *QueryGuideSubStatusResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGuideSubStatusResponseBody) SetRequestId(v string) *QueryGuideSubStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGuideSubStatusResponseBody) SetSuccess(v bool) *QueryGuideSubStatusResponseBody {
	s.Success = &v
	return s
}

func (s *QueryGuideSubStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
