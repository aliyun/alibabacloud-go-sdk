// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKMSSecretsForHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ImportKMSSecretsForHostResponseBody
	GetRequestId() *string
	SetResults(v []*ImportKMSSecretsForHostResponseBodyResults) *ImportKMSSecretsForHostResponseBody
	GetResults() []*ImportKMSSecretsForHostResponseBodyResults
	SetSuccessCount(v int64) *ImportKMSSecretsForHostResponseBody
	GetSuccessCount() *int64
}

type ImportKMSSecretsForHostResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*ImportKMSSecretsForHostResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	SuccessCount *int64 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
}

func (s ImportKMSSecretsForHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportKMSSecretsForHostResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKMSSecretsForHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportKMSSecretsForHostResponseBody) GetResults() []*ImportKMSSecretsForHostResponseBodyResults {
	return s.Results
}

func (s *ImportKMSSecretsForHostResponseBody) GetSuccessCount() *int64 {
	return s.SuccessCount
}

func (s *ImportKMSSecretsForHostResponseBody) SetRequestId(v string) *ImportKMSSecretsForHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportKMSSecretsForHostResponseBody) SetResults(v []*ImportKMSSecretsForHostResponseBodyResults) *ImportKMSSecretsForHostResponseBody {
	s.Results = v
	return s
}

func (s *ImportKMSSecretsForHostResponseBody) SetSuccessCount(v int64) *ImportKMSSecretsForHostResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *ImportKMSSecretsForHostResponseBody) Validate() error {
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

type ImportKMSSecretsForHostResponseBodyResults struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// test1
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
}

func (s ImportKMSSecretsForHostResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s ImportKMSSecretsForHostResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ImportKMSSecretsForHostResponseBodyResults) GetCode() *string {
	return s.Code
}

func (s *ImportKMSSecretsForHostResponseBodyResults) GetMessage() *string {
	return s.Message
}

func (s *ImportKMSSecretsForHostResponseBodyResults) GetSecretName() *string {
	return s.SecretName
}

func (s *ImportKMSSecretsForHostResponseBodyResults) SetCode(v string) *ImportKMSSecretsForHostResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ImportKMSSecretsForHostResponseBodyResults) SetMessage(v string) *ImportKMSSecretsForHostResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ImportKMSSecretsForHostResponseBodyResults) SetSecretName(v string) *ImportKMSSecretsForHostResponseBodyResults {
	s.SecretName = &v
	return s
}

func (s *ImportKMSSecretsForHostResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
