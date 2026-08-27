// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToGroupResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveGroupOutputFileToGroupResourceResponseBody
	GetCode() *string
	SetMessage(v string) *SaveGroupOutputFileToGroupResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveGroupOutputFileToGroupResourceResponseBody
	GetRequestId() *string
	SetResults(v []*SaveGroupOutputFileToGroupResourceResponseBodyResults) *SaveGroupOutputFileToGroupResourceResponseBody
	GetResults() []*SaveGroupOutputFileToGroupResourceResponseBodyResults
}

type SaveGroupOutputFileToGroupResourceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of results.
	Results []*SaveGroupOutputFileToGroupResourceResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s SaveGroupOutputFileToGroupResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToGroupResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToGroupResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveGroupOutputFileToGroupResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveGroupOutputFileToGroupResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveGroupOutputFileToGroupResourceResponseBody) GetResults() []*SaveGroupOutputFileToGroupResourceResponseBodyResults {
	return s.Results
}

func (s *SaveGroupOutputFileToGroupResourceResponseBody) SetCode(v string) *SaveGroupOutputFileToGroupResourceResponseBody {
	s.Code = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponseBody) SetMessage(v string) *SaveGroupOutputFileToGroupResourceResponseBody {
	s.Message = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponseBody) SetRequestId(v string) *SaveGroupOutputFileToGroupResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponseBody) SetResults(v []*SaveGroupOutputFileToGroupResourceResponseBodyResults) *SaveGroupOutputFileToGroupResourceResponseBody {
	s.Results = v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponseBody) Validate() error {
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

type SaveGroupOutputFileToGroupResourceResponseBodyResults struct {
	// The business error code (i18n key), returned when the operation fails.
	//
	// example:
	//
	// ERR.Robject.UserOutput.ItemNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error description (internationalized based on the request locale), returned when the operation fails.
	//
	// example:
	//
	// Group output does not exist
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The output item ID.
	//
	// example:
	//
	// item-1
	ItemId *string `json:"itemId,omitempty" xml:"itemId,omitempty"`
	// The sourceId of the newly created resource, returned when the operation is successful.
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

func (s SaveGroupOutputFileToGroupResourceResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToGroupResourceResponseBodyResults) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) GetItemId() *string {
	return s.ItemId
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) GetSourceId() *string {
	return s.SourceId
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) GetSuccess() *bool {
	return s.Success
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) SetErrorCode(v string) *SaveGroupOutputFileToGroupResourceResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) SetErrorMessage(v string) *SaveGroupOutputFileToGroupResourceResponseBodyResults {
	s.ErrorMessage = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) SetItemId(v string) *SaveGroupOutputFileToGroupResourceResponseBodyResults {
	s.ItemId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) SetSourceId(v string) *SaveGroupOutputFileToGroupResourceResponseBodyResults {
	s.SourceId = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) SetSuccess(v bool) *SaveGroupOutputFileToGroupResourceResponseBodyResults {
	s.Success = &v
	return s
}

func (s *SaveGroupOutputFileToGroupResourceResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
