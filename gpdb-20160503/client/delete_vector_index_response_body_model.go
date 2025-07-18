// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVectorIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteVectorIndexResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVectorIndexResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteVectorIndexResponseBody
	GetStatus() *string
}

type DeleteVectorIndexResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **success**
	//
	// 	- **fail**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteVectorIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVectorIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVectorIndexResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVectorIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVectorIndexResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteVectorIndexResponseBody) SetMessage(v string) *DeleteVectorIndexResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVectorIndexResponseBody) SetRequestId(v string) *DeleteVectorIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVectorIndexResponseBody) SetStatus(v string) *DeleteVectorIndexResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteVectorIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
