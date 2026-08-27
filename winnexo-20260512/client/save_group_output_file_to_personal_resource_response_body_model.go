// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGroupOutputFileToPersonalResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveGroupOutputFileToPersonalResourceResponseBody
	GetCode() *string
	SetMessage(v string) *SaveGroupOutputFileToPersonalResourceResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveGroupOutputFileToPersonalResourceResponseBody
	GetRequestId() *string
	SetResults(v []*SaveGroupOutputFileToPersonalResourceResponseBodyResults) *SaveGroupOutputFileToPersonalResourceResponseBody
	GetResults() []*SaveGroupOutputFileToPersonalResourceResponseBodyResults
}

type SaveGroupOutputFileToPersonalResourceResponseBody struct {
	// SUCCESS indicates success. In case of failure, the corresponding error type is returned, such as ERR_BAD_REQUEST, ERR_VALIDATION_FAILED, or ERR_INTERNAL_SERVER_ERROR.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status code description.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result list.
	Results []*SaveGroupOutputFileToPersonalResourceResponseBodyResults `json:"results,omitempty" xml:"results,omitempty" type:"Repeated"`
}

func (s SaveGroupOutputFileToPersonalResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToPersonalResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBody) GetResults() []*SaveGroupOutputFileToPersonalResourceResponseBodyResults {
	return s.Results
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBody) SetCode(v string) *SaveGroupOutputFileToPersonalResourceResponseBody {
	s.Code = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBody) SetMessage(v string) *SaveGroupOutputFileToPersonalResourceResponseBody {
	s.Message = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBody) SetRequestId(v string) *SaveGroupOutputFileToPersonalResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBody) SetResults(v []*SaveGroupOutputFileToPersonalResourceResponseBodyResults) *SaveGroupOutputFileToPersonalResourceResponseBody {
	s.Results = v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBody) Validate() error {
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

type SaveGroupOutputFileToPersonalResourceResponseBodyResults struct {
	// The business error code (i18n key), returned on failure.
	//
	// example:
	//
	// ERR.Robject.UserOutput.ItemNotFound
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error description (internationalized based on the request locale), returned on failure.
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
	// The sourceId of the newly created resource, returned on success.
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

func (s SaveGroupOutputFileToPersonalResourceResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s SaveGroupOutputFileToPersonalResourceResponseBodyResults) GoString() string {
	return s.String()
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) GetItemId() *string {
	return s.ItemId
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) GetSourceId() *string {
	return s.SourceId
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) GetSuccess() *bool {
	return s.Success
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) SetErrorCode(v string) *SaveGroupOutputFileToPersonalResourceResponseBodyResults {
	s.ErrorCode = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) SetErrorMessage(v string) *SaveGroupOutputFileToPersonalResourceResponseBodyResults {
	s.ErrorMessage = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) SetItemId(v string) *SaveGroupOutputFileToPersonalResourceResponseBodyResults {
	s.ItemId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) SetSourceId(v string) *SaveGroupOutputFileToPersonalResourceResponseBodyResults {
	s.SourceId = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) SetSuccess(v bool) *SaveGroupOutputFileToPersonalResourceResponseBodyResults {
	s.Success = &v
	return s
}

func (s *SaveGroupOutputFileToPersonalResourceResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
