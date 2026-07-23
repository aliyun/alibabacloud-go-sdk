// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteExperimentRunResponseBody
	GetRequestId() *string
}

type DeleteExperimentRunResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 019F89BE-9190-3AAA-B5A4-DBAE3BABBBEA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteExperimentRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentRunResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExperimentRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExperimentRunResponseBody) SetRequestId(v string) *DeleteExperimentRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExperimentRunResponseBody) Validate() error {
	return dara.Validate(s)
}
