// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKVCacheStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResults(v []*AttachKVCacheStoreResponseBodyAttachResults) *AttachKVCacheStoreResponseBody
	GetAttachResults() []*AttachKVCacheStoreResponseBodyAttachResults
	SetRequestId(v string) *AttachKVCacheStoreResponseBody
	GetRequestId() *string
}

type AttachKVCacheStoreResponseBody struct {
	// The list of mount results.
	AttachResults []*AttachKVCacheStoreResponseBodyAttachResults `json:"AttachResults,omitempty" xml:"AttachResults,omitempty" type:"Repeated"`
	// The request ID. A request ID is returned regardless of whether the operation is successful.
	//
	// example:
	//
	// F7BEF1E2-7E56-5BF1-8E36-77A51C5812F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachKVCacheStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachKVCacheStoreResponseBody) GoString() string {
	return s.String()
}

func (s *AttachKVCacheStoreResponseBody) GetAttachResults() []*AttachKVCacheStoreResponseBodyAttachResults {
	return s.AttachResults
}

func (s *AttachKVCacheStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachKVCacheStoreResponseBody) SetAttachResults(v []*AttachKVCacheStoreResponseBodyAttachResults) *AttachKVCacheStoreResponseBody {
	s.AttachResults = v
	return s
}

func (s *AttachKVCacheStoreResponseBody) SetRequestId(v string) *AttachKVCacheStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachKVCacheStoreResponseBody) Validate() error {
	if s.AttachResults != nil {
		for _, item := range s.AttachResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachKVCacheStoreResponseBodyAttachResults struct {
	// The error code returned when the operation fails. This value is null when the operation succeeds.
	//
	// example:
	//
	// InvalidStatus.OperationDenied
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the operation fails. This value is null when the operation succeeds.
	//
	// example:
	//
	// The operation is not allowed for instance in ATTACHED status.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// KVCacheStore KvcsId
	//
	// example:
	//
	// kvcs-xxxxx
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The operation result. Valid values:
	//
	// - ATTACHING: The request has been accepted and the asynchronous mount is in progress.
	//
	// - ATTACHED: The resource is already mounted (idempotent).
	//
	// - Success: The synchronous verification passed and the asynchronous operation is complete.
	//
	// - Failed: The operation failed.
	//
	// example:
	//
	// ATTACHING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VSC ID on the compute side.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s AttachKVCacheStoreResponseBodyAttachResults) String() string {
	return dara.Prettify(s)
}

func (s AttachKVCacheStoreResponseBodyAttachResults) GoString() string {
	return s.String()
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) GetKvcsId() *string {
	return s.KvcsId
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) GetStatus() *string {
	return s.Status
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) GetVscId() *string {
	return s.VscId
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) SetErrorCode(v string) *AttachKVCacheStoreResponseBodyAttachResults {
	s.ErrorCode = &v
	return s
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) SetErrorMessage(v string) *AttachKVCacheStoreResponseBodyAttachResults {
	s.ErrorMessage = &v
	return s
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) SetKvcsId(v string) *AttachKVCacheStoreResponseBodyAttachResults {
	s.KvcsId = &v
	return s
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) SetStatus(v string) *AttachKVCacheStoreResponseBodyAttachResults {
	s.Status = &v
	return s
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) SetVscId(v string) *AttachKVCacheStoreResponseBodyAttachResults {
	s.VscId = &v
	return s
}

func (s *AttachKVCacheStoreResponseBodyAttachResults) Validate() error {
	return dara.Validate(s)
}
