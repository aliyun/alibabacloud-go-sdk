// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DeleteRulesResponseBody
	GetJobId() *string
	SetRequestId(v string) *DeleteRulesResponseBody
	GetRequestId() *string
}

type DeleteRulesResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7BED4F62-3E6E-5E4F-8C53-2D8CCE77F2ED
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRulesResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *DeleteRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRulesResponseBody) SetJobId(v string) *DeleteRulesResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteRulesResponseBody) SetRequestId(v string) *DeleteRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
