// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCustomAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ImportCustomAlertRulesResponseBody
	GetData() *string
	SetRequestId(v string) *ImportCustomAlertRulesResponseBody
	GetRequestId() *string
}

type ImportCustomAlertRulesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportCustomAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportCustomAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportCustomAlertRulesResponseBody) GetData() *string {
	return s.Data
}

func (s *ImportCustomAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportCustomAlertRulesResponseBody) SetData(v string) *ImportCustomAlertRulesResponseBody {
	s.Data = &v
	return s
}

func (s *ImportCustomAlertRulesResponseBody) SetRequestId(v string) *ImportCustomAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportCustomAlertRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
