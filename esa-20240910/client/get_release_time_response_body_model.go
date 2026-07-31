// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReleaseTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReleaseTime(v string) *GetReleaseTimeResponseBody
	GetReleaseTime() *string
	SetRequestId(v string) *GetReleaseTimeResponseBody
	GetRequestId() *string
}

type GetReleaseTimeResponseBody struct {
	// The scheduled release time.
	//
	// example:
	//
	// 2026-01-02T06:00:00Z
	ReleaseTime *string `json:"ReleaseTime,omitempty" xml:"ReleaseTime,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6abd807e-ed2a-****-ac54-ac38a62472e6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetReleaseTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseTimeResponseBody) GoString() string {
	return s.String()
}

func (s *GetReleaseTimeResponseBody) GetReleaseTime() *string {
	return s.ReleaseTime
}

func (s *GetReleaseTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetReleaseTimeResponseBody) SetReleaseTime(v string) *GetReleaseTimeResponseBody {
	s.ReleaseTime = &v
	return s
}

func (s *GetReleaseTimeResponseBody) SetRequestId(v string) *GetReleaseTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetReleaseTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
