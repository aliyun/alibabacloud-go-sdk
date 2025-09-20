// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexFileMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *IndexFileMetaResponseBody
	GetEventId() *string
	SetRequestId(v string) *IndexFileMetaResponseBody
	GetRequestId() *string
}

type IndexFileMetaResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 30F-1D8FxFzDXKJH9YQdve4CjR****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6E93D6C9-5AC0-49F9-914D-E02678D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IndexFileMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IndexFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *IndexFileMetaResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *IndexFileMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IndexFileMetaResponseBody) SetEventId(v string) *IndexFileMetaResponseBody {
	s.EventId = &v
	return s
}

func (s *IndexFileMetaResponseBody) SetRequestId(v string) *IndexFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *IndexFileMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
