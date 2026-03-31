// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAggregateRemediationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationDeleteResults(v []*DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) *DeleteAggregateRemediationsResponseBody
	GetRemediationDeleteResults() []*DeleteAggregateRemediationsResponseBodyRemediationDeleteResults
	SetRequestId(v string) *DeleteAggregateRemediationsResponseBody
	GetRequestId() *string
}

type DeleteAggregateRemediationsResponseBody struct {
	// The results of the delete operation.
	RemediationDeleteResults []*DeleteAggregateRemediationsResponseBodyRemediationDeleteResults `json:"RemediationDeleteResults,omitempty" xml:"RemediationDeleteResults,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4BE28FB1-616A-5586-82E4-F34FB2AF7441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAggregateRemediationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateRemediationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAggregateRemediationsResponseBody) GetRemediationDeleteResults() []*DeleteAggregateRemediationsResponseBodyRemediationDeleteResults {
	return s.RemediationDeleteResults
}

func (s *DeleteAggregateRemediationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAggregateRemediationsResponseBody) SetRemediationDeleteResults(v []*DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) *DeleteAggregateRemediationsResponseBody {
	s.RemediationDeleteResults = v
	return s
}

func (s *DeleteAggregateRemediationsResponseBody) SetRequestId(v string) *DeleteAggregateRemediationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAggregateRemediationsResponseBody) Validate() error {
	if s.RemediationDeleteResults != nil {
		for _, item := range s.RemediationDeleteResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteAggregateRemediationsResponseBodyRemediationDeleteResults struct {
	// The error code returned.
	//
	// 	- If the remediation template is deleted, no error code is returned.
	//
	// 	- If the remediation template fails to be deleted, an error code is returned. For more information about error codes, see [Error codes](https://error-center.alibabacloud.com/status/product/Config).
	//
	// example:
	//
	// RemediationConfigNotExist
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the remediation template.
	//
	// example:
	//
	// crr-909ba2d4716700eb****
	RemediationId *string `json:"RemediationId,omitempty" xml:"RemediationId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) String() string {
	return dara.Prettify(s)
}

func (s DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) GoString() string {
	return s.String()
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) GetRemediationId() *string {
	return s.RemediationId
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) SetErrorMessage(v string) *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) SetRemediationId(v string) *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults {
	s.RemediationId = &v
	return s
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) SetSuccess(v bool) *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults {
	s.Success = &v
	return s
}

func (s *DeleteAggregateRemediationsResponseBodyRemediationDeleteResults) Validate() error {
	return dara.Validate(s)
}
