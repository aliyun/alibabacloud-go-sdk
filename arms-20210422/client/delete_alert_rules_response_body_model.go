// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsSuccess(v bool) *DeleteAlertRulesResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteAlertRulesResponseBody
	GetRequestId() *string
}

type DeleteAlertRulesResponseBody struct {
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlertRulesResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlertRulesResponseBody) SetIsSuccess(v bool) *DeleteAlertRulesResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteAlertRulesResponseBody) SetRequestId(v string) *DeleteAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlertRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
