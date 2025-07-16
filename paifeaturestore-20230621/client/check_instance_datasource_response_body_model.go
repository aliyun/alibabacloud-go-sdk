// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckInstanceDatasourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckInstanceDatasourceResponseBody
	GetRequestId() *string
	SetStatus(v string) *CheckInstanceDatasourceResponseBody
	GetStatus() *string
}

type CheckInstanceDatasourceResponseBody struct {
	// example:
	//
	// C03B2680-AC9C-59CD-93C5-8142B92537FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckInstanceDatasourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckInstanceDatasourceResponseBody) GoString() string {
	return s.String()
}

func (s *CheckInstanceDatasourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckInstanceDatasourceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CheckInstanceDatasourceResponseBody) SetRequestId(v string) *CheckInstanceDatasourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckInstanceDatasourceResponseBody) SetStatus(v string) *CheckInstanceDatasourceResponseBody {
	s.Status = &v
	return s
}

func (s *CheckInstanceDatasourceResponseBody) Validate() error {
	return dara.Validate(s)
}
