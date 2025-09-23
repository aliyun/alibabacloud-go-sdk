// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCronJobPolicyServerlessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyCronJobPolicyServerlessResponseBody
	GetRequestId() *string
}

type ModifyCronJobPolicyServerlessResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCronJobPolicyServerlessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCronJobPolicyServerlessResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCronJobPolicyServerlessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCronJobPolicyServerlessResponseBody) SetRequestId(v string) *ModifyCronJobPolicyServerlessResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCronJobPolicyServerlessResponseBody) Validate() error {
	return dara.Validate(s)
}
