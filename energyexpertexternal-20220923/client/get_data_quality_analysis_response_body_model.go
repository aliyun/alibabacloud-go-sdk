// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDataQualityAnalysisResponseBodyData) *GetDataQualityAnalysisResponseBody
	GetData() *GetDataQualityAnalysisResponseBodyData
	SetRequestId(v string) *GetDataQualityAnalysisResponseBody
	GetRequestId() *string
}

type GetDataQualityAnalysisResponseBody struct {
	// The response parameters.
	Data *GetDataQualityAnalysisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request. The value is unique for each request. This facilitates subsequent troubleshooting.
	//
	// example:
	//
	// 4A0AEC56-5C9A-5D47-93DF-7227836FFF82
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDataQualityAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBody) GetData() *GetDataQualityAnalysisResponseBodyData {
	return s.Data
}

func (s *GetDataQualityAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataQualityAnalysisResponseBody) SetData(v *GetDataQualityAnalysisResponseBodyData) *GetDataQualityAnalysisResponseBody {
	s.Data = v
	return s
}

func (s *GetDataQualityAnalysisResponseBody) SetRequestId(v string) *GetDataQualityAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityAnalysisResponseBodyData struct {
	// Score of each inventory.
	DataQuality []*GetDataQualityAnalysisResponseBodyDataDataQuality `json:"dataQuality,omitempty" xml:"dataQuality,omitempty" type:"Repeated"`
	// Data quality result.
	DataQualityResult *GetDataQualityAnalysisResponseBodyDataDataQualityResult `json:"dataQualityResult,omitempty" xml:"dataQualityResult,omitempty" type:"Struct"`
	// Sensitivity analysis list
	SensitivityList []*GetDataQualityAnalysisResponseBodyDataSensitivityList `json:"sensitivityList,omitempty" xml:"sensitivityList,omitempty" type:"Repeated"`
	// Uncertainty value. The model\\"s overall percentage uncertainty results. "10.00%" symbolizes a 10.00% uncertainty, indicating that the carbon footprint lies within ±10.00%. This is derived from a weighted aggregation of individual inventory uncertainties.
	//
	// example:
	//
	// 10.00
	Uncertainty *string `json:"uncertainty,omitempty" xml:"uncertainty,omitempty"`
	// Uncertainty list
	UncertaintyValues []*GetDataQualityAnalysisResponseBodyDataUncertaintyValues `json:"uncertaintyValues,omitempty" xml:"uncertaintyValues,omitempty" type:"Repeated"`
}

func (s GetDataQualityAnalysisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyData) GetDataQuality() []*GetDataQualityAnalysisResponseBodyDataDataQuality {
	return s.DataQuality
}

func (s *GetDataQualityAnalysisResponseBodyData) GetDataQualityResult() *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	return s.DataQualityResult
}

func (s *GetDataQualityAnalysisResponseBodyData) GetSensitivityList() []*GetDataQualityAnalysisResponseBodyDataSensitivityList {
	return s.SensitivityList
}

func (s *GetDataQualityAnalysisResponseBodyData) GetUncertainty() *string {
	return s.Uncertainty
}

func (s *GetDataQualityAnalysisResponseBodyData) GetUncertaintyValues() []*GetDataQualityAnalysisResponseBodyDataUncertaintyValues {
	return s.UncertaintyValues
}

func (s *GetDataQualityAnalysisResponseBodyData) SetDataQuality(v []*GetDataQualityAnalysisResponseBodyDataDataQuality) *GetDataQualityAnalysisResponseBodyData {
	s.DataQuality = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyData) SetDataQualityResult(v *GetDataQualityAnalysisResponseBodyDataDataQualityResult) *GetDataQualityAnalysisResponseBodyData {
	s.DataQualityResult = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyData) SetSensitivityList(v []*GetDataQualityAnalysisResponseBodyDataSensitivityList) *GetDataQualityAnalysisResponseBodyData {
	s.SensitivityList = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyData) SetUncertainty(v string) *GetDataQualityAnalysisResponseBodyData {
	s.Uncertainty = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyData) SetUncertaintyValues(v []*GetDataQualityAnalysisResponseBodyDataUncertaintyValues) *GetDataQualityAnalysisResponseBodyData {
	s.UncertaintyValues = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyData) Validate() error {
	if s.DataQuality != nil {
		for _, item := range s.DataQuality {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataQualityResult != nil {
		if err := s.DataQualityResult.Validate(); err != nil {
			return err
		}
	}
	if s.SensitivityList != nil {
		for _, item := range s.SensitivityList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UncertaintyValues != nil {
		for _, item := range s.UncertaintyValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataQualityAnalysisResponseBodyDataDataQuality struct {
	// Inventory name
	//
	// example:
	//
	// energy
	Inventory *string `json:"inventory,omitempty" xml:"inventory,omitempty"`
	// Score. The distribution ranges from 1 to 5. A value closer to 1 indicates better data quality.
	Score *GetDataQualityAnalysisResponseBodyDataDataQualityScore `json:"score,omitempty" xml:"score,omitempty" type:"Struct"`
}

func (s GetDataQualityAnalysisResponseBodyDataDataQuality) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataDataQuality) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQuality) GetInventory() *string {
	return s.Inventory
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQuality) GetScore() *GetDataQualityAnalysisResponseBodyDataDataQualityScore {
	return s.Score
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQuality) SetInventory(v string) *GetDataQualityAnalysisResponseBodyDataDataQuality {
	s.Inventory = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQuality) SetScore(v *GetDataQualityAnalysisResponseBodyDataDataQualityScore) *GetDataQualityAnalysisResponseBodyDataDataQuality {
	s.Score = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQuality) Validate() error {
	if s.Score != nil {
		if err := s.Score.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataQualityAnalysisResponseBodyDataDataQualityScore struct {
	// Data quality evaluation indicator 1: activity data credibility.
	//
	// example:
	//
	// 3
	G1 *float64 `json:"g1,omitempty" xml:"g1,omitempty"`
	// Data quality evaluation indicator 2: data factor reliability.
	//
	// example:
	//
	// 3
	G2 *float64 `json:"g2,omitempty" xml:"g2,omitempty"`
	// Data quality evaluation indicator 3: time representativeness.
	//
	// example:
	//
	// 3
	G3 *float64 `json:"g3,omitempty" xml:"g3,omitempty"`
	// Data quality evaluation indicator 4: geographic representativeness.
	//
	// example:
	//
	// 3
	G4 *float64 `json:"g4,omitempty" xml:"g4,omitempty"`
}

func (s GetDataQualityAnalysisResponseBodyDataDataQualityScore) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataDataQualityScore) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) GetG1() *float64 {
	return s.G1
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) GetG2() *float64 {
	return s.G2
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) GetG3() *float64 {
	return s.G3
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) GetG4() *float64 {
	return s.G4
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) SetG1(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityScore {
	s.G1 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) SetG2(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityScore {
	s.G2 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) SetG3(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityScore {
	s.G3 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) SetG4(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityScore {
	s.G4 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityScore) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityAnalysisResponseBodyDataDataQualityResult struct {
	// The score. This parameter is applicable to DQR results. The distribution ranges from 1 to 5. A value closer to 1 indicates better data quality. The number of valid digits is kept to four decimal places.
	//
	// example:
	//
	// 1.2345
	DataQualityScore *float64 `json:"data_quality_score,omitempty" xml:"data_quality_score,omitempty"`
	// Data quality evaluation indicator 1: activity data credibility.
	//
	// example:
	//
	// 1.2345
	G1 *float64 `json:"g1,omitempty" xml:"g1,omitempty"`
	// Data quality evaluation indicator 2: data factor reliability.
	//
	// example:
	//
	// 1.2345
	G2 *float64 `json:"g2,omitempty" xml:"g2,omitempty"`
	// Data quality evaluation indicator 3: time representativeness.
	//
	// example:
	//
	// 1.2345
	G3 *float64 `json:"g3,omitempty" xml:"g3,omitempty"`
	// Data quality evaluation indicator 4: geographic representativeness.
	//
	// example:
	//
	// 1.2345
	G4 *float64 `json:"g4,omitempty" xml:"g4,omitempty"`
}

func (s GetDataQualityAnalysisResponseBodyDataDataQualityResult) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataDataQualityResult) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) GetDataQualityScore() *float64 {
	return s.DataQualityScore
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) GetG1() *float64 {
	return s.G1
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) GetG2() *float64 {
	return s.G2
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) GetG3() *float64 {
	return s.G3
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) GetG4() *float64 {
	return s.G4
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetDataQualityScore(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.DataQualityScore = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetG1(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.G1 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetG2(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.G2 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetG3(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.G3 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) SetG4(v float64) *GetDataQualityAnalysisResponseBodyDataDataQualityResult {
	s.G4 = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataDataQualityResult) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityAnalysisResponseBodyDataSensitivityList struct {
	// Inventory id
	//
	// example:
	//
	// 1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Name of the inventory item.
	//
	// example:
	//
	// energy
	Inventory *string `json:"inventory,omitempty" xml:"inventory,omitempty"`
	// List of emission reduction measures.
	ReductionList []*string `json:"reductionList,omitempty" xml:"reductionList,omitempty" type:"Repeated"`
	// Sensitivity percentage.
	//
	// example:
	//
	// 91.7
	Sensitivity *float64 `json:"sensitivity,omitempty" xml:"sensitivity,omitempty"`
}

func (s GetDataQualityAnalysisResponseBodyDataSensitivityList) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataSensitivityList) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) GetId() *string {
	return s.Id
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) GetInventory() *string {
	return s.Inventory
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) GetReductionList() []*string {
	return s.ReductionList
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) GetSensitivity() *float64 {
	return s.Sensitivity
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) SetId(v string) *GetDataQualityAnalysisResponseBodyDataSensitivityList {
	s.Id = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) SetInventory(v string) *GetDataQualityAnalysisResponseBodyDataSensitivityList {
	s.Inventory = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) SetReductionList(v []*string) *GetDataQualityAnalysisResponseBodyDataSensitivityList {
	s.ReductionList = v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) SetSensitivity(v float64) *GetDataQualityAnalysisResponseBodyDataSensitivityList {
	s.Sensitivity = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataSensitivityList) Validate() error {
	return dara.Validate(s)
}

type GetDataQualityAnalysisResponseBodyDataUncertaintyValues struct {
	// The name of the inventory. Format: process name / inventory name.
	//
	// example:
	//
	// process-1/inventory-1
	Inventory *string `json:"inventory,omitempty" xml:"inventory,omitempty"`
	// Inventory uncertainty absolute contribution size. The impact of the quality of each inventory data on the carbon footprint results in the modeling process, when the uncertain contribution of a list is large, please improve its data quality as much as possible to improve the accuracy of carbon footprint analysis. The meaning of "1.4964" is not determined to contribute 1.4964 kgCO₂ e/unit, where unit is the unit of the product.
	//
	// example:
	//
	// 1.4964
	UncertaintyContribution *string `json:"uncertaintyContribution,omitempty" xml:"uncertaintyContribution,omitempty"`
}

func (s GetDataQualityAnalysisResponseBodyDataUncertaintyValues) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityAnalysisResponseBodyDataUncertaintyValues) GoString() string {
	return s.String()
}

func (s *GetDataQualityAnalysisResponseBodyDataUncertaintyValues) GetInventory() *string {
	return s.Inventory
}

func (s *GetDataQualityAnalysisResponseBodyDataUncertaintyValues) GetUncertaintyContribution() *string {
	return s.UncertaintyContribution
}

func (s *GetDataQualityAnalysisResponseBodyDataUncertaintyValues) SetInventory(v string) *GetDataQualityAnalysisResponseBodyDataUncertaintyValues {
	s.Inventory = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataUncertaintyValues) SetUncertaintyContribution(v string) *GetDataQualityAnalysisResponseBodyDataUncertaintyValues {
	s.UncertaintyContribution = &v
	return s
}

func (s *GetDataQualityAnalysisResponseBodyDataUncertaintyValues) Validate() error {
	return dara.Validate(s)
}
