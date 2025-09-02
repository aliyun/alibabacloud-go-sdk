// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScanSensitiveDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ScanSensitiveDataResponseBody
	GetRequestId() *string
	SetSensitives(v map[string]interface{}) *ScanSensitiveDataResponseBody
	GetSensitives() map[string]interface{}
}

type ScanSensitiveDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The check result. sensDatas indicates the rules that are used to check the sensitive data. sensDatas includes the following parameters:
	//
	// 	- hitCount: the number of times that the sensitive data hits the rule.
	//
	// 	- ruleName: the name of the rule.
	Sensitives map[string]interface{} `json:"Sensitives,omitempty" xml:"Sensitives,omitempty"`
}

func (s ScanSensitiveDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScanSensitiveDataResponseBody) GoString() string {
	return s.String()
}

func (s *ScanSensitiveDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScanSensitiveDataResponseBody) GetSensitives() map[string]interface{} {
	return s.Sensitives
}

func (s *ScanSensitiveDataResponseBody) SetRequestId(v string) *ScanSensitiveDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScanSensitiveDataResponseBody) SetSensitives(v map[string]interface{}) *ScanSensitiveDataResponseBody {
	s.Sensitives = v
	return s
}

func (s *ScanSensitiveDataResponseBody) Validate() error {
	return dara.Validate(s)
}
