// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocFunctionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFunctions(v []*DescribePocFunctionsResponseBodyFunctions) *DescribePocFunctionsResponseBody
	GetFunctions() []*DescribePocFunctionsResponseBodyFunctions
	SetRequestId(v string) *DescribePocFunctionsResponseBody
	GetRequestId() *string
}

type DescribePocFunctionsResponseBody struct {
	// A list of the POC feature details.
	Functions []*DescribePocFunctionsResponseBodyFunctions `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1557B42F-B889-460A-B17F-1DE5C5AD7FF2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePocFunctionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePocFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePocFunctionsResponseBody) GetFunctions() []*DescribePocFunctionsResponseBodyFunctions {
	return s.Functions
}

func (s *DescribePocFunctionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePocFunctionsResponseBody) SetFunctions(v []*DescribePocFunctionsResponseBodyFunctions) *DescribePocFunctionsResponseBody {
	s.Functions = v
	return s
}

func (s *DescribePocFunctionsResponseBody) SetRequestId(v string) *DescribePocFunctionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePocFunctionsResponseBody) Validate() error {
	if s.Functions != nil {
		for _, item := range s.Functions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePocFunctionsResponseBodyFunctions struct {
	// The expiration time of the POC feature trial. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1760581677000
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The type of the POC feature trial. Valid values:
	//
	// - **apisec**: API security.
	//
	// - **botWeb**: bot management for websites.
	//
	// - **botApp**: bot management for apps.
	//
	// - **largeLanguageModel**: AI-powered application protection.
	//
	// example:
	//
	// botWeb
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribePocFunctionsResponseBodyFunctions) String() string {
	return dara.Prettify(s)
}

func (s DescribePocFunctionsResponseBodyFunctions) GoString() string {
	return s.String()
}

func (s *DescribePocFunctionsResponseBodyFunctions) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *DescribePocFunctionsResponseBodyFunctions) GetType() *string {
	return s.Type
}

func (s *DescribePocFunctionsResponseBodyFunctions) SetExpireTime(v int64) *DescribePocFunctionsResponseBodyFunctions {
	s.ExpireTime = &v
	return s
}

func (s *DescribePocFunctionsResponseBodyFunctions) SetType(v string) *DescribePocFunctionsResponseBodyFunctions {
	s.Type = &v
	return s
}

func (s *DescribePocFunctionsResponseBodyFunctions) Validate() error {
	return dara.Validate(s)
}
