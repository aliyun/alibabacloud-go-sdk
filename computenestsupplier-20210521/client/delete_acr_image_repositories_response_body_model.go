// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAcrImageRepositoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAcrImageRepositoriesResponseBody
	GetRequestId() *string
}

type DeleteAcrImageRepositoriesResponseBody struct {
	// example:
	//
	// 9B55A3FD-B562-5BFE-A91A-DB1790717236
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAcrImageRepositoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAcrImageRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAcrImageRepositoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAcrImageRepositoriesResponseBody) SetRequestId(v string) *DeleteAcrImageRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAcrImageRepositoriesResponseBody) Validate() error {
	return dara.Validate(s)
}
