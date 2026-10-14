// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgenticCatalogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AgenticCatalog) *GetAgenticCatalogResponseBody
	GetData() *AgenticCatalog
	SetErrorCode(v string) *GetAgenticCatalogResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetAgenticCatalogResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetAgenticCatalogResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgenticCatalogResponseBody
	GetSuccess() *bool
}

type GetAgenticCatalogResponseBody struct {
	// The catalog details.
	Data *AgenticCatalog `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned when an exception occurs.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgenticCatalogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgenticCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgenticCatalogResponseBody) GetData() *AgenticCatalog {
	return s.Data
}

func (s *GetAgenticCatalogResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetAgenticCatalogResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetAgenticCatalogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgenticCatalogResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgenticCatalogResponseBody) SetData(v *AgenticCatalog) *GetAgenticCatalogResponseBody {
	s.Data = v
	return s
}

func (s *GetAgenticCatalogResponseBody) SetErrorCode(v string) *GetAgenticCatalogResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAgenticCatalogResponseBody) SetErrorMessage(v string) *GetAgenticCatalogResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAgenticCatalogResponseBody) SetRequestId(v string) *GetAgenticCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgenticCatalogResponseBody) SetSuccess(v bool) *GetAgenticCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgenticCatalogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
