// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomAuthConnectBlackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryCustomAuthConnectBlackResponseBody
	GetCode() *int32
	SetData(v *QueryCustomAuthConnectBlackResponseBodyData) *QueryCustomAuthConnectBlackResponseBody
	GetData() *QueryCustomAuthConnectBlackResponseBodyData
	SetMessage(v string) *QueryCustomAuthConnectBlackResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCustomAuthConnectBlackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCustomAuthConnectBlackResponseBody
	GetSuccess() *bool
}

type QueryCustomAuthConnectBlackResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request is successful. Other status codes indicate that the request failed. For a list of error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryCustomAuthConnectBlackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 020F6A43-19E6-4B6E-B846-44EB31DF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCustomAuthConnectBlackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthConnectBlackResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthConnectBlackResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryCustomAuthConnectBlackResponseBody) GetData() *QueryCustomAuthConnectBlackResponseBodyData {
	return s.Data
}

func (s *QueryCustomAuthConnectBlackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCustomAuthConnectBlackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCustomAuthConnectBlackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCustomAuthConnectBlackResponseBody) SetCode(v int32) *QueryCustomAuthConnectBlackResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBody) SetData(v *QueryCustomAuthConnectBlackResponseBodyData) *QueryCustomAuthConnectBlackResponseBody {
	s.Data = v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBody) SetMessage(v string) *QueryCustomAuthConnectBlackResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBody) SetRequestId(v string) *QueryCustomAuthConnectBlackResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBody) SetSuccess(v bool) *QueryCustomAuthConnectBlackResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCustomAuthConnectBlackResponseBodyData struct {
	// The pagination token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAZ0cM0HTqLXvgm7oMHWXcvc=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The returned results.
	Results []*QueryCustomAuthConnectBlackResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s QueryCustomAuthConnectBlackResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthConnectBlackResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthConnectBlackResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCustomAuthConnectBlackResponseBodyData) GetResults() []*QueryCustomAuthConnectBlackResponseBodyDataResults {
	return s.Results
}

func (s *QueryCustomAuthConnectBlackResponseBodyData) SetNextToken(v string) *QueryCustomAuthConnectBlackResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBodyData) SetResults(v []*QueryCustomAuthConnectBlackResponseBodyDataResults) *QueryCustomAuthConnectBlackResponseBodyData {
	s.Results = v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBodyData) Validate() error {
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

type QueryCustomAuthConnectBlackResponseBodyDataResults struct {
	// The client ID.
	//
	// example:
	//
	// GID_TEST@@@test
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// Indicates whether to allow or deny access.
	//
	// example:
	//
	// ALLOW
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The authorized permissions.
	//
	// example:
	//
	// CONNECT
	PermitAction *string `json:"PermitAction,omitempty" xml:"PermitAction,omitempty"`
}

func (s QueryCustomAuthConnectBlackResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthConnectBlackResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthConnectBlackResponseBodyDataResults) GetClientId() *string {
	return s.ClientId
}

func (s *QueryCustomAuthConnectBlackResponseBodyDataResults) GetEffect() *string {
	return s.Effect
}

func (s *QueryCustomAuthConnectBlackResponseBodyDataResults) GetPermitAction() *string {
	return s.PermitAction
}

func (s *QueryCustomAuthConnectBlackResponseBodyDataResults) SetClientId(v string) *QueryCustomAuthConnectBlackResponseBodyDataResults {
	s.ClientId = &v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBodyDataResults) SetEffect(v string) *QueryCustomAuthConnectBlackResponseBodyDataResults {
	s.Effect = &v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBodyDataResults) SetPermitAction(v string) *QueryCustomAuthConnectBlackResponseBodyDataResults {
	s.PermitAction = &v
	return s
}

func (s *QueryCustomAuthConnectBlackResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
