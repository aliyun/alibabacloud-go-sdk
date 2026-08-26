// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeLiveStreamStateResponseBody
	GetRequestId() *string
	SetStreamState(v string) *DescribeLiveStreamStateResponseBody
	GetStreamState() *string
	SetType(v string) *DescribeLiveStreamStateResponseBody
	GetType() *string
}

type DescribeLiveStreamStateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CE6CD79D-0A98-1F22-A15F-FADA74DF2729
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the stream. Valid values:
	//
	// - online: The stream is active.
	//
	// - offline: The stream is offline. This may mean the stream ingest has failed or ended. For specific details, use the data returned by the stream ingest callback. This operation does not provide a detailed breakdown of the offline status.
	//
	// example:
	//
	// online
	StreamState *string `json:"StreamState,omitempty" xml:"StreamState,omitempty"`
	// The stream ingest method. Valid values:
	//
	// - push
	//
	// - pull
	//
	// example:
	//
	// push
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeLiveStreamStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamStateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamStateResponseBody) GetStreamState() *string {
	return s.StreamState
}

func (s *DescribeLiveStreamStateResponseBody) GetType() *string {
	return s.Type
}

func (s *DescribeLiveStreamStateResponseBody) SetRequestId(v string) *DescribeLiveStreamStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamStateResponseBody) SetStreamState(v string) *DescribeLiveStreamStateResponseBody {
	s.StreamState = &v
	return s
}

func (s *DescribeLiveStreamStateResponseBody) SetType(v string) *DescribeLiveStreamStateResponseBody {
	s.Type = &v
	return s
}

func (s *DescribeLiveStreamStateResponseBody) Validate() error {
	return dara.Validate(s)
}
