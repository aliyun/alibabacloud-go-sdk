// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteArtifactLifecycleRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteArtifactLifecycleRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteArtifactLifecycleRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteArtifactLifecycleRuleResponseBody
	GetRequestId() *string
}

type DeleteArtifactLifecycleRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 001AB638-C99B-5A27-8AC9-B2DBABFFEBB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteArtifactLifecycleRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteArtifactLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteArtifactLifecycleRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteArtifactLifecycleRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteArtifactLifecycleRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteArtifactLifecycleRuleResponseBody) SetCode(v string) *DeleteArtifactLifecycleRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponseBody) SetIsSuccess(v bool) *DeleteArtifactLifecycleRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponseBody) SetRequestId(v string) *DeleteArtifactLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteArtifactLifecycleRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
