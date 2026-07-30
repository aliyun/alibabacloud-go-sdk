// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *UpdateMediaResponseBody
	GetMediaId() *string
	SetRequestId(v string) *UpdateMediaResponseBody
	GetRequestId() *string
}

type UpdateMediaResponseBody struct {
	// The ICE media asset ID.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *UpdateMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaResponseBody) SetMediaId(v string) *UpdateMediaResponseBody {
	s.MediaId = &v
	return s
}

func (s *UpdateMediaResponseBody) SetRequestId(v string) *UpdateMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
