// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DeleteLiveStreamBlockResponseBody
	GetDescription() *string
	SetRequestId(v string) *DeleteLiveStreamBlockResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteLiveStreamBlockResponseBody
	GetStatus() *string
}

type DeleteLiveStreamBlockResponseBody struct {
	// The result description. If the request was successful, ok is returned. If the request failed, the failure detail is returned.
	//
	// example:
	//
	// ok
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B908FF89-B03C-4831-B55B-48D2A7DA0A68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status. Valid values:
	//
	// 	- ok: The request was successful.
	//
	// 	- fail: The request failed.
	//
	// >  If any parameter failed to be configured, the request failed.
	//
	// example:
	//
	// ok
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteLiveStreamBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamBlockResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamBlockResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DeleteLiveStreamBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveStreamBlockResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteLiveStreamBlockResponseBody) SetDescription(v string) *DeleteLiveStreamBlockResponseBody {
	s.Description = &v
	return s
}

func (s *DeleteLiveStreamBlockResponseBody) SetRequestId(v string) *DeleteLiveStreamBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveStreamBlockResponseBody) SetStatus(v string) *DeleteLiveStreamBlockResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteLiveStreamBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
