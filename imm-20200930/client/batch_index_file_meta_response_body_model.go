// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchIndexFileMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEventId(v string) *BatchIndexFileMetaResponseBody
	GetEventId() *string
	SetRequestId(v string) *BatchIndexFileMetaResponseBody
	GetRequestId() *string
}

type BatchIndexFileMetaResponseBody struct {
	// The event ID.
	//
	// example:
	//
	// 387-1DAPFFZplUZhuCuhnB6I9H****
	EventId *string `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8F93E6D9-5AC0-49F9-914D-E02678A3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchIndexFileMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchIndexFileMetaResponseBody) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaResponseBody) GetEventId() *string {
	return s.EventId
}

func (s *BatchIndexFileMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchIndexFileMetaResponseBody) SetEventId(v string) *BatchIndexFileMetaResponseBody {
	s.EventId = &v
	return s
}

func (s *BatchIndexFileMetaResponseBody) SetRequestId(v string) *BatchIndexFileMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchIndexFileMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
