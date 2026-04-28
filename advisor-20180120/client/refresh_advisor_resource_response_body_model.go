// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *RefreshAdvisorResourceResponseBody
	GetData() *int64
	SetRequestId(v string) *RefreshAdvisorResourceResponseBody
	GetRequestId() *string
}

type RefreshAdvisorResourceResponseBody struct {
	// example:
	//
	// 12345678
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 566331F9-5AB3-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RefreshAdvisorResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorResourceResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorResourceResponseBody) GetData() *int64 {
	return s.Data
}

func (s *RefreshAdvisorResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshAdvisorResourceResponseBody) SetData(v int64) *RefreshAdvisorResourceResponseBody {
	s.Data = &v
	return s
}

func (s *RefreshAdvisorResourceResponseBody) SetRequestId(v string) *RefreshAdvisorResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshAdvisorResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
