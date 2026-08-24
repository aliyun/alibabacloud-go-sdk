// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKVCacheStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetachResults(v []*DetachKVCacheStoreResponseBodyDetachResults) *DetachKVCacheStoreResponseBody
	GetDetachResults() []*DetachKVCacheStoreResponseBodyDetachResults
	SetRequestId(v string) *DetachKVCacheStoreResponseBody
	GetRequestId() *string
}

type DetachKVCacheStoreResponseBody struct {
	// The list of unmount results.
	DetachResults []*DetachKVCacheStoreResponseBodyDetachResults `json:"DetachResults,omitempty" xml:"DetachResults,omitempty" type:"Repeated"`
	// The request ID. A request ID is returned regardless of whether the API call succeeds.
	//
	// example:
	//
	// B127704C-ECB1-5B0A-AA9C-8F394A6F179F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachKVCacheStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachKVCacheStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DetachKVCacheStoreResponseBody) GetDetachResults() []*DetachKVCacheStoreResponseBodyDetachResults {
	return s.DetachResults
}

func (s *DetachKVCacheStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachKVCacheStoreResponseBody) SetDetachResults(v []*DetachKVCacheStoreResponseBodyDetachResults) *DetachKVCacheStoreResponseBody {
	s.DetachResults = v
	return s
}

func (s *DetachKVCacheStoreResponseBody) SetRequestId(v string) *DetachKVCacheStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachKVCacheStoreResponseBody) Validate() error {
	if s.DetachResults != nil {
		for _, item := range s.DetachResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetachKVCacheStoreResponseBodyDetachResults struct {
	// The error code when the operation fails. This value is null when the operation succeeds.
	//
	// example:
	//
	// KVCacheInstance.NotAttached
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message when the operation fails. This value is null when the operation succeeds.
	//
	// example:
	//
	// The KVCacheInstance is not in ATTACHED status.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// KVCacheStore KvcsId
	//
	// example:
	//
	// kvcs-xxxxx
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The operation result. Valid values:
	//
	// - DETACHING: The request has been accepted and the asynchronous unmount is in progress. This value is also returned for idempotent calls.
	//
	// - Success: The synchronous validation passed and the asynchronous operation completed.
	//
	// - Failed: The operation failed.
	//
	// example:
	//
	// DETACHING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The VSC ID on the compute side.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DetachKVCacheStoreResponseBodyDetachResults) String() string {
	return dara.Prettify(s)
}

func (s DetachKVCacheStoreResponseBodyDetachResults) GoString() string {
	return s.String()
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) GetKvcsId() *string {
	return s.KvcsId
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) GetStatus() *string {
	return s.Status
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) GetVscId() *string {
	return s.VscId
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) SetErrorCode(v string) *DetachKVCacheStoreResponseBodyDetachResults {
	s.ErrorCode = &v
	return s
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) SetErrorMessage(v string) *DetachKVCacheStoreResponseBodyDetachResults {
	s.ErrorMessage = &v
	return s
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) SetKvcsId(v string) *DetachKVCacheStoreResponseBodyDetachResults {
	s.KvcsId = &v
	return s
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) SetStatus(v string) *DetachKVCacheStoreResponseBodyDetachResults {
	s.Status = &v
	return s
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) SetVscId(v string) *DetachKVCacheStoreResponseBodyDetachResults {
	s.VscId = &v
	return s
}

func (s *DetachKVCacheStoreResponseBodyDetachResults) Validate() error {
	return dara.Validate(s)
}
