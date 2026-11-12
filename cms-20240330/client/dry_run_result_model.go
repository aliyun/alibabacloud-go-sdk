// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunResult interface {
	dara.Model
	String() string
	GoString() string
	SetExitCode(v int32) *DryRunResult
	GetExitCode() *int32
	SetResult(v string) *DryRunResult
	GetResult() *string
	SetViolations(v []*DryRunViolation) *DryRunResult
	GetViolations() []*DryRunViolation
}

type DryRunResult struct {
	// The exit code. A value of 0 indicates that the validation passed. A value of 1 indicates that the request was rejected by business rules. For specific reasons, see violations. A value of 2 indicates that the request structure cannot be parsed and rule validation cannot be performed.
	//
	// example:
	//
	// 1
	ExitCode *int32 `json:"exitCode,omitempty" xml:"exitCode,omitempty"`
	// The text representation of the result. A value of ok indicates that the validation passed, and rejected indicates that the request was rejected by business rules. This field is empty when the exit code is 2, which means the request structure cannot be parsed.
	//
	// example:
	//
	// rejected
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// The list of violations. Each violation consists of three elements: field, message, and suggestion, which represent the path of the violated field, the violation description, and the remediation suggestion, respectively. This array is empty when the exit code is 0 (passed). All violations are listed when the exit code is 1 (rejected by business rules). The location and cause of the unparseable fields are provided when the exit code is 2 (unparseable request structure).
	//
	// example:
	//
	// [{"field":"schedule.intervalSecs","message":"schedule interval must be a multiple of 15 seconds","suggestion":"must satisfy: 15-86400 seconds and a multiple of 15"}]
	Violations []*DryRunViolation `json:"violations,omitempty" xml:"violations,omitempty" type:"Repeated"`
}

func (s DryRunResult) String() string {
	return dara.Prettify(s)
}

func (s DryRunResult) GoString() string {
	return s.String()
}

func (s *DryRunResult) GetExitCode() *int32 {
	return s.ExitCode
}

func (s *DryRunResult) GetResult() *string {
	return s.Result
}

func (s *DryRunResult) GetViolations() []*DryRunViolation {
	return s.Violations
}

func (s *DryRunResult) SetExitCode(v int32) *DryRunResult {
	s.ExitCode = &v
	return s
}

func (s *DryRunResult) SetResult(v string) *DryRunResult {
	s.Result = &v
	return s
}

func (s *DryRunResult) SetViolations(v []*DryRunViolation) *DryRunResult {
	s.Violations = v
	return s
}

func (s *DryRunResult) Validate() error {
	if s.Violations != nil {
		for _, item := range s.Violations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
