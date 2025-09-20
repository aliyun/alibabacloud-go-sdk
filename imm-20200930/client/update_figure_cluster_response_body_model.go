// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFigureClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateFigureClusterResponseBody
	GetRequestId() *string
}

type UpdateFigureClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5F74C5C9-5AC0-49F9-914D-E01589D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateFigureClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateFigureClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFigureClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateFigureClusterResponseBody) SetRequestId(v string) *UpdateFigureClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFigureClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
