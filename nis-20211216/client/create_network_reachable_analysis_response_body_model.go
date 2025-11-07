// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNetworkReachableAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkReachableAnalysisId(v string) *CreateNetworkReachableAnalysisResponseBody
	GetNetworkReachableAnalysisId() *string
	SetRequestId(v string) *CreateNetworkReachableAnalysisResponseBody
	GetRequestId() *string
}

type CreateNetworkReachableAnalysisResponseBody struct {
	// The ID of the task for analyzing network reachability.
	//
	// example:
	//
	// nra-2fede05617494417****
	NetworkReachableAnalysisId *string `json:"NetworkReachableAnalysisId,omitempty" xml:"NetworkReachableAnalysisId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNetworkReachableAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNetworkReachableAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetworkReachableAnalysisResponseBody) GetNetworkReachableAnalysisId() *string {
	return s.NetworkReachableAnalysisId
}

func (s *CreateNetworkReachableAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetworkReachableAnalysisResponseBody) SetNetworkReachableAnalysisId(v string) *CreateNetworkReachableAnalysisResponseBody {
	s.NetworkReachableAnalysisId = &v
	return s
}

func (s *CreateNetworkReachableAnalysisResponseBody) SetRequestId(v string) *CreateNetworkReachableAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetworkReachableAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
