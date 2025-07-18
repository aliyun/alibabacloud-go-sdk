// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUpsertCollectionDataJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CancelUpsertCollectionDataJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelUpsertCollectionDataJobResponseBody
	GetRequestId() *string
	SetStatus(v string) *CancelUpsertCollectionDataJobResponseBody
	GetStatus() *string
}

type CancelUpsertCollectionDataJobResponseBody struct {
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

func (s CancelUpsertCollectionDataJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelUpsertCollectionDataJobResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUpsertCollectionDataJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelUpsertCollectionDataJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelUpsertCollectionDataJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CancelUpsertCollectionDataJobResponseBody) SetMessage(v string) *CancelUpsertCollectionDataJobResponseBody {
	s.Message = &v
	return s
}

func (s *CancelUpsertCollectionDataJobResponseBody) SetRequestId(v string) *CancelUpsertCollectionDataJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelUpsertCollectionDataJobResponseBody) SetStatus(v string) *CancelUpsertCollectionDataJobResponseBody {
	s.Status = &v
	return s
}

func (s *CancelUpsertCollectionDataJobResponseBody) Validate() error {
	return dara.Validate(s)
}
