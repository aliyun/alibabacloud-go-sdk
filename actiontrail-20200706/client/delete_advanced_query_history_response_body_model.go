// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAdvancedQueryHistoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAdvancedQueryHistoryResponseBody
	GetRequestId() *string
}

type DeleteAdvancedQueryHistoryResponseBody struct {
	// example:
	//
	// 04857D99-8B0C-53EB-85F1-E64198E7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAdvancedQueryHistoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAdvancedQueryHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAdvancedQueryHistoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAdvancedQueryHistoryResponseBody) SetRequestId(v string) *DeleteAdvancedQueryHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAdvancedQueryHistoryResponseBody) Validate() error {
	return dara.Validate(s)
}
