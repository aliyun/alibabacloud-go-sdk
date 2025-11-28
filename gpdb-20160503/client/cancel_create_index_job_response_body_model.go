// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCreateIndexJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CancelCreateIndexJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelCreateIndexJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *CancelCreateIndexJobResponseBody
	GetStatus() *string
}

type CancelCreateIndexJobResponseBody struct {
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
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

func (s CancelCreateIndexJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCreateIndexJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCreateIndexJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelCreateIndexJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCreateIndexJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CancelCreateIndexJobResponseBody) SetMessage(v string) *CancelCreateIndexJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelCreateIndexJobResponseBody) SetRequestId(v string) *CancelCreateIndexJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCreateIndexJobResponseBody) SetStatus(v string) *CancelCreateIndexJobResponseBody {
	s.Status = &v
	return s
}

func (s *CancelCreateIndexJobResponseBody) Validate() error {
	return dara.Validate(s)
}
