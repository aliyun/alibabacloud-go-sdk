// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRemediationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRemediationDeleteResults(v []*DeleteRemediationsResponseBodyRemediationDeleteResults) *DeleteRemediationsResponseBody
	GetRemediationDeleteResults() []*DeleteRemediationsResponseBodyRemediationDeleteResults
	SetRequestId(v string) *DeleteRemediationsResponseBody
	GetRequestId() *string
}

type DeleteRemediationsResponseBody struct {
	// The returned result.
	RemediationDeleteResults []*DeleteRemediationsResponseBodyRemediationDeleteResults `json:"RemediationDeleteResults,omitempty" xml:"RemediationDeleteResults,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 4BE28FB1-616A-5586-82E4-F34FB2AF7441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRemediationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemediationsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRemediationsResponseBody) GetRemediationDeleteResults() []*DeleteRemediationsResponseBodyRemediationDeleteResults {
	return s.RemediationDeleteResults
}

func (s *DeleteRemediationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRemediationsResponseBody) SetRemediationDeleteResults(v []*DeleteRemediationsResponseBodyRemediationDeleteResults) *DeleteRemediationsResponseBody {
	s.RemediationDeleteResults = v
	return s
}

func (s *DeleteRemediationsResponseBody) SetRequestId(v string) *DeleteRemediationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRemediationsResponseBody) Validate() error {
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

type DeleteRemediationsResponseBodyRemediationDeleteResults struct {
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

func (s DeleteRemediationsResponseBodyRemediationDeleteResults) String() string {
	return dara.Prettify(s)
}

func (s DeleteRemediationsResponseBodyRemediationDeleteResults) GoString() string {
	return s.String()
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) GetRemediationId() *string {
	return s.RemediationId
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) SetErrorMessage(v string) *DeleteRemediationsResponseBodyRemediationDeleteResults {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) SetRemediationId(v string) *DeleteRemediationsResponseBodyRemediationDeleteResults {
	s.RemediationId = &v
	return s
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) SetSuccess(v bool) *DeleteRemediationsResponseBodyRemediationDeleteResults {
	s.Success = &v
	return s
}

func (s *DeleteRemediationsResponseBodyRemediationDeleteResults) Validate() error {
	return dara.Validate(s)
}
