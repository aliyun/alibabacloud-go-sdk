// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveRealtimeLogAuthorizedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedStatus(v string) *DescribeLiveRealtimeLogAuthorizedResponseBody
	GetAuthorizedStatus() *string
	SetRequestId(v string) *DescribeLiveRealtimeLogAuthorizedResponseBody
	GetRequestId() *string
}

type DescribeLiveRealtimeLogAuthorizedResponseBody struct {
	// The authorization status. **true**: authorized **false**: not authorized
	//
	// example:
	//
	// true
	AuthorizedStatus *string `json:"AuthorizedStatus,omitempty" xml:"AuthorizedStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91b******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveRealtimeLogAuthorizedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveRealtimeLogAuthorizedResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveRealtimeLogAuthorizedResponseBody) GetAuthorizedStatus() *string {
	return s.AuthorizedStatus
}

func (s *DescribeLiveRealtimeLogAuthorizedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveRealtimeLogAuthorizedResponseBody) SetAuthorizedStatus(v string) *DescribeLiveRealtimeLogAuthorizedResponseBody {
	s.AuthorizedStatus = &v
	return s
}

func (s *DescribeLiveRealtimeLogAuthorizedResponseBody) SetRequestId(v string) *DescribeLiveRealtimeLogAuthorizedResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveRealtimeLogAuthorizedResponseBody) Validate() error {
	return dara.Validate(s)
}
