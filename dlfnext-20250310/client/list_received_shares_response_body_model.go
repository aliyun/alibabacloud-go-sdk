// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReceivedSharesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextPageToken(v string) *ListReceivedSharesResponseBody
	GetNextPageToken() *string
	SetShares(v []*ReceivedShare) *ListReceivedSharesResponseBody
	GetShares() []*ReceivedShare
}

type ListReceivedSharesResponseBody struct {
	// example:
	//
	// ""
	NextPageToken *string          `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
	Shares        []*ReceivedShare `json:"shares,omitempty" xml:"shares,omitempty" type:"Repeated"`
}

func (s ListReceivedSharesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListReceivedSharesResponseBody) GoString() string {
	return s.String()
}

func (s *ListReceivedSharesResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListReceivedSharesResponseBody) GetShares() []*ReceivedShare {
	return s.Shares
}

func (s *ListReceivedSharesResponseBody) SetNextPageToken(v string) *ListReceivedSharesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListReceivedSharesResponseBody) SetShares(v []*ReceivedShare) *ListReceivedSharesResponseBody {
	s.Shares = v
	return s
}

func (s *ListReceivedSharesResponseBody) Validate() error {
	return dara.Validate(s)
}
