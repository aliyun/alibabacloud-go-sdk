// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteServiceLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceLabelResponseBody
	GetRequestId() *string
}

type DeleteServiceLabelResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Succeed to delete service [service_from_zxxx] labels.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceLabelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceLabelResponseBody) SetMessage(v string) *DeleteServiceLabelResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceLabelResponseBody) SetRequestId(v string) *DeleteServiceLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
