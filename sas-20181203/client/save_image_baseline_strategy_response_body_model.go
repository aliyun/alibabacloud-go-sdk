// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveImageBaselineStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveImageBaselineStrategyResponseBody
	GetRequestId() *string
}

type SaveImageBaselineStrategyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9AB5D3DE-6E0F-5633-AA71-4B90C724****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SaveImageBaselineStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveImageBaselineStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *SaveImageBaselineStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveImageBaselineStrategyResponseBody) SetRequestId(v string) *SaveImageBaselineStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveImageBaselineStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
