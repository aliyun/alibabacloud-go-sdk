// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataProcessJobTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataProcessJobTemplateResponseBody
	GetRequestId() *string
}

type DeleteDataProcessJobTemplateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataProcessJobTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataProcessJobTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataProcessJobTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataProcessJobTemplateResponseBody) SetRequestId(v string) *DeleteDataProcessJobTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataProcessJobTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
