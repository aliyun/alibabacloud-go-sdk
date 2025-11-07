// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkReachableAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteNetworkReachableAnalysisResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteNetworkReachableAnalysisResponseBody
	GetRequestId() *string
}

type DeleteNetworkReachableAnalysisResponseBody struct {
	// Result of operation.
	//
	// - **true**: Delete Success.
	//
	// - **false**: Delete Fail.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4838F3F2-30E1-5D82-B25A-B9FE33BC3E25
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkReachableAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkReachableAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkReachableAnalysisResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteNetworkReachableAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkReachableAnalysisResponseBody) SetData(v bool) *DeleteNetworkReachableAnalysisResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNetworkReachableAnalysisResponseBody) SetRequestId(v string) *DeleteNetworkReachableAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkReachableAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
