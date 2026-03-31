// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateRemediationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAggregatorId(v string) *DeleteAggregateRemediationsRequest
	GetAggregatorId() *string
	SetRemediationIds(v string) *DeleteAggregateRemediationsRequest
	GetRemediationIds() *string
}

type DeleteAggregateRemediationsRequest struct {
	// The ID of the account group.
	//
	// For more information about how to obtain the ID of the account group, see [ListAggregators](https://help.aliyun.com/document_detail/255797.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ca-6b4a626622af0012****
	AggregatorId *string `json:"AggregatorId,omitempty" xml:"AggregatorId,omitempty"`
	// The ID of the remediation template. Separate multiple remediation template IDs with commas (,).
	//
	// For more information about how to obtain the ID of a remediation template, see [ListAggregateRemediations](https://help.aliyun.com/document_detail/270036.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// crr-909ba2d4716700eb****
	RemediationIds *string `json:"RemediationIds,omitempty" xml:"RemediationIds,omitempty"`
}

func (s DeleteAggregateRemediationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateRemediationsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAggregateRemediationsRequest) GetAggregatorId() *string {
	return s.AggregatorId
}

func (s *DeleteAggregateRemediationsRequest) GetRemediationIds() *string {
	return s.RemediationIds
}

func (s *DeleteAggregateRemediationsRequest) SetAggregatorId(v string) *DeleteAggregateRemediationsRequest {
	s.AggregatorId = &v
	return s
}

func (s *DeleteAggregateRemediationsRequest) SetRemediationIds(v string) *DeleteAggregateRemediationsRequest {
	s.RemediationIds = &v
	return s
}

func (s *DeleteAggregateRemediationsRequest) Validate() error {
	return dara.Validate(s)
}
