// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineCodeDeploymentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCodeVersions(v []*CreateRoutineCodeDeploymentResponseBodyCodeVersions) *CreateRoutineCodeDeploymentResponseBody
	GetCodeVersions() []*CreateRoutineCodeDeploymentResponseBodyCodeVersions
	SetDeploymentId(v string) *CreateRoutineCodeDeploymentResponseBody
	GetDeploymentId() *string
	SetRequestId(v string) *CreateRoutineCodeDeploymentResponseBody
	GetRequestId() *string
	SetStrategy(v string) *CreateRoutineCodeDeploymentResponseBody
	GetStrategy() *string
}

type CreateRoutineCodeDeploymentResponseBody struct {
	// The configuration list of the phased release version number.
	CodeVersions []*CreateRoutineCodeDeploymentResponseBodyCodeVersions `json:"CodeVersions,omitempty" xml:"CodeVersions,omitempty" type:"Repeated"`
	// The deployment record ID.
	//
	// example:
	//
	// 234
	DeploymentId *string `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The phased release policy. The constant string is "percentage".
	//
	// example:
	//
	// percentage
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s CreateRoutineCodeDeploymentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineCodeDeploymentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoutineCodeDeploymentResponseBody) GetCodeVersions() []*CreateRoutineCodeDeploymentResponseBodyCodeVersions {
	return s.CodeVersions
}

func (s *CreateRoutineCodeDeploymentResponseBody) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *CreateRoutineCodeDeploymentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoutineCodeDeploymentResponseBody) GetStrategy() *string {
	return s.Strategy
}

func (s *CreateRoutineCodeDeploymentResponseBody) SetCodeVersions(v []*CreateRoutineCodeDeploymentResponseBodyCodeVersions) *CreateRoutineCodeDeploymentResponseBody {
	s.CodeVersions = v
	return s
}

func (s *CreateRoutineCodeDeploymentResponseBody) SetDeploymentId(v string) *CreateRoutineCodeDeploymentResponseBody {
	s.DeploymentId = &v
	return s
}

func (s *CreateRoutineCodeDeploymentResponseBody) SetRequestId(v string) *CreateRoutineCodeDeploymentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoutineCodeDeploymentResponseBody) SetStrategy(v string) *CreateRoutineCodeDeploymentResponseBody {
	s.Strategy = &v
	return s
}

func (s *CreateRoutineCodeDeploymentResponseBody) Validate() error {
	if s.CodeVersions != nil {
		for _, item := range s.CodeVersions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateRoutineCodeDeploymentResponseBodyCodeVersions struct {
	// The version of the code.
	//
	// example:
	//
	// 1723599747213377175
	CodeVersion *string `json:"CodeVersion,omitempty" xml:"CodeVersion,omitempty"`
	// The phased release ratio.
	//
	// example:
	//
	// 100
	Percentage *int64 `json:"Percentage,omitempty" xml:"Percentage,omitempty"`
}

func (s CreateRoutineCodeDeploymentResponseBodyCodeVersions) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineCodeDeploymentResponseBodyCodeVersions) GoString() string {
	return s.String()
}

func (s *CreateRoutineCodeDeploymentResponseBodyCodeVersions) GetCodeVersion() *string {
	return s.CodeVersion
}

func (s *CreateRoutineCodeDeploymentResponseBodyCodeVersions) GetPercentage() *int64 {
	return s.Percentage
}

func (s *CreateRoutineCodeDeploymentResponseBodyCodeVersions) SetCodeVersion(v string) *CreateRoutineCodeDeploymentResponseBodyCodeVersions {
	s.CodeVersion = &v
	return s
}

func (s *CreateRoutineCodeDeploymentResponseBodyCodeVersions) SetPercentage(v int64) *CreateRoutineCodeDeploymentResponseBodyCodeVersions {
	s.Percentage = &v
	return s
}

func (s *CreateRoutineCodeDeploymentResponseBodyCodeVersions) Validate() error {
	return dara.Validate(s)
}
