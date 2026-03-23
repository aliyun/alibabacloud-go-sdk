// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyApplicationPromptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ApplyApplicationPromptsResponseBody
	GetApplicationId() *string
	SetRequestId(v string) *ApplyApplicationPromptsResponseBody
	GetRequestId() *string
}

type ApplyApplicationPromptsResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApplyApplicationPromptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyApplicationPromptsResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyApplicationPromptsResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ApplyApplicationPromptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyApplicationPromptsResponseBody) SetApplicationId(v string) *ApplyApplicationPromptsResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *ApplyApplicationPromptsResponseBody) SetRequestId(v string) *ApplyApplicationPromptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyApplicationPromptsResponseBody) Validate() error {
	return dara.Validate(s)
}
