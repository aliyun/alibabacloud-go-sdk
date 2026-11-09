// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCustomAuthPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryCustomAuthPermissionResponseBody
	GetCode() *int32
	SetData(v *QueryCustomAuthPermissionResponseBodyData) *QueryCustomAuthPermissionResponseBody
	GetData() *QueryCustomAuthPermissionResponseBodyData
	SetMessage(v string) *QueryCustomAuthPermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryCustomAuthPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryCustomAuthPermissionResponseBody
	GetSuccess() *bool
}

type QueryCustomAuthPermissionResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryCustomAuthPermissionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 63309FDB-ED6C-46AE-B31C-A172FBA0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryCustomAuthPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthPermissionResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryCustomAuthPermissionResponseBody) GetData() *QueryCustomAuthPermissionResponseBodyData {
	return s.Data
}

func (s *QueryCustomAuthPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryCustomAuthPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryCustomAuthPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryCustomAuthPermissionResponseBody) SetCode(v int32) *QueryCustomAuthPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBody) SetData(v *QueryCustomAuthPermissionResponseBodyData) *QueryCustomAuthPermissionResponseBody {
	s.Data = v
	return s
}

func (s *QueryCustomAuthPermissionResponseBody) SetMessage(v string) *QueryCustomAuthPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBody) SetRequestId(v string) *QueryCustomAuthPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBody) SetSuccess(v bool) *QueryCustomAuthPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryCustomAuthPermissionResponseBodyData struct {
	// The token that marks the end position of the previous returned page. To obtain the next batch of data, call the operation again by using the value of NextToken returned by the previous request. If you call this operation for the first time or want to query all results, set NextToken to an empty string.
	//
	// example:
	//
	// AAAAAV/vsqTyeMlX1MIk7/b6NrZLIlsSVf49O04ac7HAmlBoaYspakK7ZZkR3vRDp5Y9Nz0EmuWYrtF+1qkUwuJzPk/qEto/FGxl5Kd+qdwNt3t8
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The response results.
	Results []*QueryCustomAuthPermissionResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s QueryCustomAuthPermissionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthPermissionResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthPermissionResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *QueryCustomAuthPermissionResponseBodyData) GetResults() []*QueryCustomAuthPermissionResponseBodyDataResults {
	return s.Results
}

func (s *QueryCustomAuthPermissionResponseBodyData) SetNextToken(v string) *QueryCustomAuthPermissionResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBodyData) SetResults(v []*QueryCustomAuthPermissionResponseBodyDataResults) *QueryCustomAuthPermissionResponseBodyData {
	s.Results = v
	return s
}

func (s *QueryCustomAuthPermissionResponseBodyData) Validate() error {
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

type QueryCustomAuthPermissionResponseBodyDataResults struct {
	// Indicates whether to allow or deny access.
	//
	// example:
	//
	// ALLOW
	Effect *string `json:"Effect,omitempty" xml:"Effect,omitempty"`
	// The username or client ID.
	//
	// example:
	//
	// test
	Identity *string `json:"Identity,omitempty" xml:"Identity,omitempty"`
	// The identity type. Valid values:
	//
	// 	- USER
	//
	// 	- CLIENT
	//
	// example:
	//
	// USER
	IdentityType *string `json:"IdentityType,omitempty" xml:"IdentityType,omitempty"`
	// The authorized permissions.
	//
	// example:
	//
	// PUB_SUB
	PermitAction *string `json:"PermitAction,omitempty" xml:"PermitAction,omitempty"`
	// The topic name. Multi-level topics and wildcard characters are supported.
	//
	// example:
	//
	// test
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s QueryCustomAuthPermissionResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s QueryCustomAuthPermissionResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) GetEffect() *string {
	return s.Effect
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) GetIdentity() *string {
	return s.Identity
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) GetIdentityType() *string {
	return s.IdentityType
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) GetPermitAction() *string {
	return s.PermitAction
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) GetTopic() *string {
	return s.Topic
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) SetEffect(v string) *QueryCustomAuthPermissionResponseBodyDataResults {
	s.Effect = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) SetIdentity(v string) *QueryCustomAuthPermissionResponseBodyDataResults {
	s.Identity = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) SetIdentityType(v string) *QueryCustomAuthPermissionResponseBodyDataResults {
	s.IdentityType = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) SetPermitAction(v string) *QueryCustomAuthPermissionResponseBodyDataResults {
	s.PermitAction = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) SetTopic(v string) *QueryCustomAuthPermissionResponseBodyDataResults {
	s.Topic = &v
	return s
}

func (s *QueryCustomAuthPermissionResponseBodyDataResults) Validate() error {
	return dara.Validate(s)
}
