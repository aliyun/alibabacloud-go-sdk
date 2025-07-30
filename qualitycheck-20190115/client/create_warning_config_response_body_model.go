// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarningConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWarningConfigResponseBody
	GetCode() *string
	SetData(v string) *CreateWarningConfigResponseBody
	GetData() *string
	SetMessage(v string) *CreateWarningConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWarningConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWarningConfigResponseBody
	GetSuccess() *bool
}

type CreateWarningConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 30
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWarningConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWarningConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarningConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWarningConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateWarningConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWarningConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWarningConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWarningConfigResponseBody) SetCode(v string) *CreateWarningConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetData(v string) *CreateWarningConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetMessage(v string) *CreateWarningConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetRequestId(v string) *CreateWarningConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarningConfigResponseBody) SetSuccess(v bool) *CreateWarningConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWarningConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
