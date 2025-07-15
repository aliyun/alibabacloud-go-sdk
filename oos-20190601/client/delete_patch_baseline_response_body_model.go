// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePatchBaselineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePatchBaselineResponseBody
	GetRequestId() *string
}

type DeletePatchBaselineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 85197066-0209-5775-BBED-99DF9DA44E48
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePatchBaselineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePatchBaselineResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePatchBaselineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePatchBaselineResponseBody) SetRequestId(v string) *DeletePatchBaselineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePatchBaselineResponseBody) Validate() error {
	return dara.Validate(s)
}
