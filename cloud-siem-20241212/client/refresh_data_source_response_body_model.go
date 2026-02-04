// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RefreshDataSourceResponseBody
	GetRequestId() *string
}

type RefreshDataSourceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshDataSourceResponseBody) SetRequestId(v string) *RefreshDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
