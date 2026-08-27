// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOutputFileToResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveOutputFileToResourceResponseBody
	GetCode() *string
	SetMessage(v string) *SaveOutputFileToResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveOutputFileToResourceResponseBody
	GetRequestId() *string
	SetResults(v []*SaveOutputFileToResourceResponseBodyResults) *SaveOutputFileToResourceResponseBody
	GetResults() []*SaveOutputFileToResourceResponseBodyResults
}

type SaveOutputFileToResourceResponseBody struct {
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- / InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description. This value is empty on success.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The per-record results in the same order as the input itemIds. A single record failure does not affect other records.
	Results []*SaveOutputFileToResourceResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s SaveOutputFileToResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveOutputFileToResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveOutputFileToResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveOutputFileToResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveOutputFileToResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveOutputFileToResourceResponseBody) GetResults() []*SaveOutputFileToResourceResponseBodyResults {
	return s.Results
}

func (s *SaveOutputFileToResourceResponseBody) SetCode(v string) *SaveOutputFileToResourceResponseBody {
	s.Code = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBody) SetMessage(v string) *SaveOutputFileToResourceResponseBody {
	s.Message = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBody) SetRequestId(v string) *SaveOutputFileToResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBody) SetResults(v []*SaveOutputFileToResourceResponseBodyResults) *SaveOutputFileToResourceResponseBody {
	s.Results = v
	return s
}

func (s *SaveOutputFileToResourceResponseBody) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveOutputFileToResourceResponseBodyResults struct {
	// The business error code (i18n key). Returned on failure.
	//
	// example:
	//
	// string_value
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error description, localized based on the request Accept-Language header. Returned on failure.
	//
	// example:
	//
	// string_value
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The output detail ID.
	//
	// example:
	//
	// exampleItemId
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The sourceId of the newly created resource. Returned on success.
	//
	// example:
	//
	// exampleSourceId
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s SaveOutputFileToResourceResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s SaveOutputFileToResourceResponseBodyResults) GoString() string {
	return s.String()
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetItemId() *string {
	return s.ItemId
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetSourceId() *string {
	return s.SourceId
}

func (s *SaveOutputFileToResourceResponseBodyResults) GetSuccess() *bool {
	return s.Success
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetErrorCode(v string) *SaveOutputFileToResourceResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetErrorMessage(v string) *SaveOutputFileToResourceResponseBodyResults {
	s.ErrorMessage = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetItemId(v string) *SaveOutputFileToResourceResponseBodyResults {
	s.ItemId = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetSourceId(v string) *SaveOutputFileToResourceResponseBodyResults {
	s.SourceId = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) SetSuccess(v bool) *SaveOutputFileToResourceResponseBodyResults {
	s.Success = &v
	return s
}

func (s *SaveOutputFileToResourceResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
