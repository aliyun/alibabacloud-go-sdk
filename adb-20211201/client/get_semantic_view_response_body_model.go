// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSemanticViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SemanticViewModel) *GetSemanticViewResponseBody
	GetData() *SemanticViewModel
	SetRequestId(v string) *GetSemanticViewResponseBody
	GetRequestId() *string
}

type GetSemanticViewResponseBody struct {
	// The details of the semantic view.
	//
	// example:
	//
	// 69
	Data *SemanticViewModel `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSemanticViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSemanticViewResponseBody) GoString() string {
	return s.String()
}

func (s *GetSemanticViewResponseBody) GetData() *SemanticViewModel {
	return s.Data
}

func (s *GetSemanticViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSemanticViewResponseBody) SetData(v *SemanticViewModel) *GetSemanticViewResponseBody {
	s.Data = v
	return s
}

func (s *GetSemanticViewResponseBody) SetRequestId(v string) *GetSemanticViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSemanticViewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
