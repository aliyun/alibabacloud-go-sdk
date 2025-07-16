// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteResourceLogResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteResourceLogResponseBody
	GetRequestId() *string
}

type DeleteResourceLogResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Log service for resource [eas-r-asdasdasd] is deleting
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteResourceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteResourceLogResponseBody) SetMessage(v string) *DeleteResourceLogResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceLogResponseBody) SetRequestId(v string) *DeleteResourceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteResourceLogResponseBody) Validate() error {
	return dara.Validate(s)
}
