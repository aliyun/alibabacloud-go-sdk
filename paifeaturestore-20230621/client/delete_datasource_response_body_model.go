// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatasourceResponseBody
	GetRequestId() *string
}

type DeleteDatasourceResponseBody struct {
	// example:
	//
	// E2E1575F-29D1-5579-B649-B7883A793562
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasourceResponseBody) SetRequestId(v string) *DeleteDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}
