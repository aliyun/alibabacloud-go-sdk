// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpsertCollectionDataAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *UpsertCollectionDataAsyncResponseBody
	GetJobId() *string
	SetMessage(v string) *UpsertCollectionDataAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpsertCollectionDataAsyncResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpsertCollectionDataAsyncResponseBody
	GetStatus() *string
}

type UpsertCollectionDataAsyncResponseBody struct {
	// The job ID. It can be used to query the job status or cancel the job.
	//
	// example:
	//
	// 231460f8-75dc-405e-a669-0c5204887e91
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
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

func (s UpsertCollectionDataAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpsertCollectionDataAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *UpsertCollectionDataAsyncResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpsertCollectionDataAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpsertCollectionDataAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpsertCollectionDataAsyncResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpsertCollectionDataAsyncResponseBody) SetJobId(v string) *UpsertCollectionDataAsyncResponseBody {
	s.JobId = &v
	return s
}

func (s *UpsertCollectionDataAsyncResponseBody) SetMessage(v string) *UpsertCollectionDataAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *UpsertCollectionDataAsyncResponseBody) SetRequestId(v string) *UpsertCollectionDataAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpsertCollectionDataAsyncResponseBody) SetStatus(v string) *UpsertCollectionDataAsyncResponseBody {
	s.Status = &v
	return s
}

func (s *UpsertCollectionDataAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
