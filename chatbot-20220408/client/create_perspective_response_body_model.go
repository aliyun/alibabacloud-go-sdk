// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePerspectiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPerspectiveId(v string) *CreatePerspectiveResponseBody
	GetPerspectiveId() *string
	SetRequestId(v string) *CreatePerspectiveResponseBody
	GetRequestId() *string
}

type CreatePerspectiveResponseBody struct {
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	// example:
	//
	// F285D735-D580-18A8-B97F-B2E72B00F101
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePerspectiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePerspectiveResponseBody) GetPerspectiveId() *string {
	return s.PerspectiveId
}

func (s *CreatePerspectiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePerspectiveResponseBody) SetPerspectiveId(v string) *CreatePerspectiveResponseBody {
	s.PerspectiveId = &v
	return s
}

func (s *CreatePerspectiveResponseBody) SetRequestId(v string) *CreatePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePerspectiveResponseBody) Validate() error {
	return dara.Validate(s)
}
