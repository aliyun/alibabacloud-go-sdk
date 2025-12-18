// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutEvaluationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteMode(v bool) *PutEvaluationsRequest
	GetDeleteMode() *bool
	SetEvaluations(v string) *PutEvaluationsRequest
	GetEvaluations() *string
	SetResultToken(v string) *PutEvaluationsRequest
	GetResultToken() *string
}

type PutEvaluationsRequest struct {
	// Specifies whether to enable the delete mode. Valid values:
	//
	// 	- true: enables the delete mode
	//
	// 	- false (default): disables the delete mode
	//
	// > This parameter is valid only when you manually trigger or periodically trigger custom rules to evaluate resources. If you enable the delete mode, the evaluation results that are not updated during the current evaluation are automatically deleted.
	//
	// example:
	//
	// false
	DeleteMode *bool `json:"DeleteMode,omitempty" xml:"DeleteMode,omitempty"`
	// The evaluation results.
	//
	// example:
	//
	// [{"accountId":120886317861****,"annotation":"The flow log is not enabled.","complianceResourceId":"flowlog-o6wdfo1yvgo4i8****","complianceResourceType":"ACS::CEN::Flowlog","complianceRegionId":"cn-shanghai","complianceType":"NON_COMPLIANT","orderingTimestamp":1588907220408}]
	Evaluations *string `json:"Evaluations,omitempty" xml:"Evaluations,omitempty"`
	// The callback token. When Cloud Config triggers a custom rule to evaluate resources, the token information is sent to Function Compute as an input parameter. The token must be specified when you submit the evaluation results.
	//
	// This parameter is required.
	//
	// example:
	//
	// =lAUbfkWp7GL9AFoQEIStinqBMc4FC8sHvip/1F1npkWUDNS2GEm6xwL6Zl/fSr0bbkWY+aiCLjTJxnp4H/yp/8p/Q8VCAtqG5uhRii4sfnYRnTPnE****
	ResultToken *string `json:"ResultToken,omitempty" xml:"ResultToken,omitempty"`
}

func (s PutEvaluationsRequest) String() string {
	return dara.Prettify(s)
}

func (s PutEvaluationsRequest) GoString() string {
	return s.String()
}

func (s *PutEvaluationsRequest) GetDeleteMode() *bool {
	return s.DeleteMode
}

func (s *PutEvaluationsRequest) GetEvaluations() *string {
	return s.Evaluations
}

func (s *PutEvaluationsRequest) GetResultToken() *string {
	return s.ResultToken
}

func (s *PutEvaluationsRequest) SetDeleteMode(v bool) *PutEvaluationsRequest {
	s.DeleteMode = &v
	return s
}

func (s *PutEvaluationsRequest) SetEvaluations(v string) *PutEvaluationsRequest {
	s.Evaluations = &v
	return s
}

func (s *PutEvaluationsRequest) SetResultToken(v string) *PutEvaluationsRequest {
	s.ResultToken = &v
	return s
}

func (s *PutEvaluationsRequest) Validate() error {
	return dara.Validate(s)
}
