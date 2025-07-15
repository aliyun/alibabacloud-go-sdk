// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayChoosenShowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PlayChoosenShowResponseBody
	GetRequestId() *string
	SetShowId(v string) *PlayChoosenShowResponseBody
	GetShowId() *string
}

type PlayChoosenShowResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the episode.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
}

func (s PlayChoosenShowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PlayChoosenShowResponseBody) GoString() string {
	return s.String()
}

func (s *PlayChoosenShowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PlayChoosenShowResponseBody) GetShowId() *string {
	return s.ShowId
}

func (s *PlayChoosenShowResponseBody) SetRequestId(v string) *PlayChoosenShowResponseBody {
	s.RequestId = &v
	return s
}

func (s *PlayChoosenShowResponseBody) SetShowId(v string) *PlayChoosenShowResponseBody {
	s.ShowId = &v
	return s
}

func (s *PlayChoosenShowResponseBody) Validate() error {
	return dara.Validate(s)
}
