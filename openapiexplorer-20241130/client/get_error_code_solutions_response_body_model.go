// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorCodeSolutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetErrorCodeSolutionsResponseBody
	GetRequestId() *string
	SetSolutions(v []*GetErrorCodeSolutionsResponseBodySolutions) *GetErrorCodeSolutionsResponseBody
	GetSolutions() []*GetErrorCodeSolutionsResponseBodySolutions
}

type GetErrorCodeSolutionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// A707AFA8-1A4C-5B2A-A165-8436C1EA38DB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The solutions. Not all error codes have corresponding solutions. You can submit a ticket or use OpenAPI Explorer to contact technical support if necessary.
	Solutions []*GetErrorCodeSolutionsResponseBodySolutions `json:"solutions,omitempty" xml:"solutions,omitempty" type:"Repeated"`
}

func (s GetErrorCodeSolutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetErrorCodeSolutionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetErrorCodeSolutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetErrorCodeSolutionsResponseBody) GetSolutions() []*GetErrorCodeSolutionsResponseBodySolutions {
	return s.Solutions
}

func (s *GetErrorCodeSolutionsResponseBody) SetRequestId(v string) *GetErrorCodeSolutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBody) SetSolutions(v []*GetErrorCodeSolutionsResponseBodySolutions) *GetErrorCodeSolutionsResponseBody {
	s.Solutions = v
	return s
}

func (s *GetErrorCodeSolutionsResponseBody) Validate() error {
	if s.Solutions != nil {
		for _, item := range s.Solutions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetErrorCodeSolutionsResponseBodySolutions struct {
	// The content of the solution, which is in the markdown format.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The error code.
	//
	// example:
	//
	// 0017-00000502
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The resource is not in a valid state for the operation.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The code of the product to which the solution belongs.
	//
	// example:
	//
	// Ecs
	Product *string `json:"product,omitempty" xml:"product,omitempty"`
	// The name of the product to which the solution belongs.
	ProductName *string `json:"productName,omitempty" xml:"productName,omitempty"`
	// The solution ID.
	//
	// example:
	//
	// 0017-00000502
	SolutionId *string `json:"solutionId,omitempty" xml:"solutionId,omitempty"`
}

func (s GetErrorCodeSolutionsResponseBodySolutions) String() string {
	return dara.Prettify(s)
}

func (s GetErrorCodeSolutionsResponseBodySolutions) GoString() string {
	return s.String()
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) GetContent() *string {
	return s.Content
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) GetProduct() *string {
	return s.Product
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) GetProductName() *string {
	return s.ProductName
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) GetSolutionId() *string {
	return s.SolutionId
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetContent(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.Content = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetErrorCode(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.ErrorCode = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetErrorMessage(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.ErrorMessage = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetProduct(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.Product = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetProductName(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.ProductName = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) SetSolutionId(v string) *GetErrorCodeSolutionsResponseBodySolutions {
	s.SolutionId = &v
	return s
}

func (s *GetErrorCodeSolutionsResponseBodySolutions) Validate() error {
	return dara.Validate(s)
}
