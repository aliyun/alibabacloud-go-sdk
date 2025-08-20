// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisDimensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientIps(v []*string) *DescribeDiagnosisDimensionsResponseBody
	GetClientIps() []*string
	SetDatabases(v []*string) *DescribeDiagnosisDimensionsResponseBody
	GetDatabases() []*string
	SetRequestId(v string) *DescribeDiagnosisDimensionsResponseBody
	GetRequestId() *string
	SetResourceGroups(v []*string) *DescribeDiagnosisDimensionsResponseBody
	GetResourceGroups() []*string
	SetUserNames(v []*string) *DescribeDiagnosisDimensionsResponseBody
	GetUserNames() []*string
}

type DescribeDiagnosisDimensionsResponseBody struct {
	// The queried source IP addresses.
	ClientIps []*string `json:"ClientIps,omitempty" xml:"ClientIps,omitempty" type:"Repeated"`
	// The queried database names.
	Databases []*string `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DEA97C6B-D7A4-5E69-9EFC-D7F88737CED5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried resource group names.
	ResourceGroups []*string `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// The queried usernames.
	UserNames []*string `json:"UserNames,omitempty" xml:"UserNames,omitempty" type:"Repeated"`
}

func (s DescribeDiagnosisDimensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisDimensionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisDimensionsResponseBody) GetClientIps() []*string {
	return s.ClientIps
}

func (s *DescribeDiagnosisDimensionsResponseBody) GetDatabases() []*string {
	return s.Databases
}

func (s *DescribeDiagnosisDimensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisDimensionsResponseBody) GetResourceGroups() []*string {
	return s.ResourceGroups
}

func (s *DescribeDiagnosisDimensionsResponseBody) GetUserNames() []*string {
	return s.UserNames
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetClientIps(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.ClientIps = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetDatabases(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetRequestId(v string) *DescribeDiagnosisDimensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetResourceGroups(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) SetUserNames(v []*string) *DescribeDiagnosisDimensionsResponseBody {
	s.UserNames = v
	return s
}

func (s *DescribeDiagnosisDimensionsResponseBody) Validate() error {
	return dara.Validate(s)
}
