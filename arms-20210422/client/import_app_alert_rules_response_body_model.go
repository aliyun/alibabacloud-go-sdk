// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAppAlertRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *ImportAppAlertRulesResponseBody
	GetData() *string
	SetRequestId(v string) *ImportAppAlertRulesResponseBody
	GetRequestId() *string
}

type ImportAppAlertRulesResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportAppAlertRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportAppAlertRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesResponseBody) GetData() *string {
	return s.Data
}

func (s *ImportAppAlertRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportAppAlertRulesResponseBody) SetData(v string) *ImportAppAlertRulesResponseBody {
	s.Data = &v
	return s
}

func (s *ImportAppAlertRulesResponseBody) SetRequestId(v string) *ImportAppAlertRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportAppAlertRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
