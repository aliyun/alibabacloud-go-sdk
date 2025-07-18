// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteIndexResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIndexResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteIndexResponseBody
	GetStatus() *string
}

type DeleteIndexResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIndexResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteIndexResponseBody) SetMessage(v string) *DeleteIndexResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIndexResponseBody) SetRequestId(v string) *DeleteIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndexResponseBody) SetStatus(v string) *DeleteIndexResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
