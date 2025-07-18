// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteCollectionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteCollectionResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteCollectionResponseBody
	GetStatus() *string
}

type DeleteCollectionResponseBody struct {
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

func (s DeleteCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCollectionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCollectionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteCollectionResponseBody) SetMessage(v string) *DeleteCollectionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCollectionResponseBody) SetRequestId(v string) *DeleteCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCollectionResponseBody) SetStatus(v string) *DeleteCollectionResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
